package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mohdjishin/GoCart/database"
	I "github.com/mohdjishin/GoCart/interfaces"
	"github.com/mohdjishin/GoCart/model"
	utils "github.com/mohdjishin/GoCart/utils"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct{}

var DB = database.NewDatabaseConnection()

func NewAdmin() I.IAdmin {
	return &Admin{}
}

func (*Admin) Signup(c *fiber.Ctx) error {
	db := DB.OpenDb()

	defer DB.CloseDb(db)
	// get the username and password
	user := new(model.Admin)

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "not able to hash pasword",
		})

	}

	usr := model.Admin{Username: user.Username, Password: string(hash), Email: user.Email}

	// create the user

	res := db.Create(&usr)

	if res.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create account",
		})

	}
	return c.Status(200).SendString("account created")
}

func (*Admin) Login(c *fiber.Ctx) error {
	db := DB.OpenDb()
	defer DB.CloseDb(db)

	body := new(model.Admin)
	// take data from req
	if err := c.BodyParser(body); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	err := utils.ValidateStruct(body)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	usr := new(model.Admin)
	db.First(&usr, "username=?", body.Username)
	if usr.ID == 0 {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid  usermame or password",
		})
	}

	// check pass

	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(body.Password))
	if err != nil {

		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "incorrect  password!",
		})
	}

	fmt.Println(usr.ID)
	// refresh token
	uuidv4, _ := uuid.NewRandom()

	err = db.Model(&usr).Where("id = ?", usr.ID).Update("refresh", uuidv4.String()).Error
	if err != nil {
		fmt.Println(err)

	}

	tokenString, errMessage := Token.GenJwtToken("admin", usr.ID, 86400)
	if errMessage != "" {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": errMessage,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"access_token":  tokenString,
		"refresh_token": uuidv4,
	})

}

func (*Admin) Validate(c *fiber.Ctx) error {

	user := c.Locals("id")
	fmt.Println(user)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "login success",
	})

}
func (*Admin) UserManagement(c *fiber.Ctx) error {
	db := DB.OpenDb()
	defer DB.CloseDb(db)

	user := new(model.Users)

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	fmt.Println(user.ID)

	fmt.Println(user.Status)

	db.Model(&model.Users{}).Where("id = ?", user.ID).Update("status", user.Status)
	return c.Status(200).JSON(fiber.Map{
		"message": "updated",
	})

}
func (*Admin) ViewUsers(c *fiber.Ctx) error {

	db := DB.OpenDb()
	defer DB.CloseDb(db)
	var user []model.Users
	err := db.Where("id > ?", 0).Find(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	var address []model.Address
	err = db.Where("id > ?", 0).Find(&address).Error
	if err != nil {
		fmt.Println(err)
	}

	infoWithoutAddress := utils.ExtractPersonalInfo(user)
	infowithPersonal := utils.ExtractAdresses(address)

	comb := utils.Combined(infoWithoutAddress, infowithPersonal)

	return c.Status(200).JSON(comb)

}

func (*Admin) Logout(c *fiber.Ctx) error {
	userId := c.Locals("id")
	usr_id := fmt.Sprintf("%v", userId)

	u64, err := strconv.ParseUint(usr_id, 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	tokenString, errMessage := Token.GenJwtToken("admin", uint(u64), 1)
	if errMessage != "" {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": errMessage,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"access_tokem": tokenString,
	})

}

func (*Admin) ViewOrders(c *fiber.Ctx) error {

	db := DB.OpenDb()
	defer DB.CloseDb(db)

	var orders []model.Order

	db.Find(&orders)

	extracttedOrderInfo := utils.ExtractOrderInfo(orders)

	return c.Status(200).JSON(extracttedOrderInfo)

}

func (*Admin) DeliveryStatusUpdate(c *fiber.Ctx) error {
	db := DB.OpenDb()
	defer DB.CloseDb(db)
	type delStatus struct {
		Id              int    `json:"id"`
		ProID           int    `json:"pro_id"`
		ShippmentStatus string `json:"shipment_status"`
	}

	order := new(model.Order)

	st := new(delStatus)

	if err := c.BodyParser(st); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	if st.ShippmentStatus == "processing" || st.ShippmentStatus == "delivered" || st.ShippmentStatus == "shipped" || st.ShippmentStatus == "deleyed" || st.ShippmentStatus == "close" {
		db.First(&order, "id = ? and product_id= ?", st.Id, st.ProID)

		if order.ID == 0 {
			return c.Status(404).JSON(fiber.Map{
				"message": "not fount",
			})
		}

		if st.ShippmentStatus == "close" {
			db.Model(&order).Where("id = ? and product_id= ?", st.Id, st.ProID).Update("shippment_status", "closed")

			db.Delete(&order, "id = ? and product_id= ?", st.Id, st.ProID)
			return c.Status(200).JSON(fiber.Map{
				"message": "order closed",
			})
		}

		db.Model(&order).Where("id = ? and product_id= ?", st.Id, st.ProID).Update("shippment_status", st.ShippmentStatus)

		return c.Status(200).JSON(fiber.Map{
			"meaage": "status updated",
		})
	}

	// processing

	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
		"message": "shippment status error",
	})

}
func (*Admin) ManageUser(c *fiber.Ctx) error {
	db := DB.OpenDb()
	defer DB.CloseDb(db)
	user := new(model.Users)

	type manageUser struct {
		ID    int  `json:"id"`
		Block bool `json:"block"`
	}
	u := new(manageUser)
	if err := c.BodyParser(u); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	db.First(&user, u.ID)

	user.Blocked = u.Block
	db.Save(&user)
	return c.Status(200).JSON(user)

}

func (*Admin) AdminRefresh(c *fiber.Ctx) error {
	db := DB.OpenDb()
	defer DB.CloseDb(db)
	type refreshToken struct {
		Access_token  string `json:"access_token"`
		Refresh_token string `json:"refresh_token"`
	}
	rt := new(refreshToken)

	if err := c.BodyParser(rt); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	err, accessToken, rfToken := Token.AdminRefreshToken(db, rt.Refresh_token, rt.Access_token)

	if err != "" {
		return c.Status(400).JSON(fiber.Map{"message": err})
	}
	return c.Status(200).JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": rfToken,
	})
}
