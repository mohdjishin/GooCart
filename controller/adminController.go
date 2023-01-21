package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mohdjishin/GoCart/database"
	"github.com/mohdjishin/GoCart/model"
	utils "github.com/mohdjishin/GoCart/utils"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {
	db := database.OpenDb()

	database.CloseDb(db)
	// get the username and password
	user := new(model.Admin)

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	fmt.Println(user.Username)
	fmt.Println(user.Email)
	fmt.Println(user.Password)
	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "not able to hash pasword",
		})

	}
	fmt.Println(hash)

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

func Login(c *fiber.Ctx) error {
	db := database.OpenDb()
	defer database.CloseDb(db)
	fmt.Println("hh")

	body := new(model.Admin)
	// take data from req
	if err := c.BodyParser(body); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	usr := new(model.Admin)
	db.First(&usr, "username=?", body.Username)
	if usr.ID == 0 {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid  usermame or password",
		})
	}

	// check pass
	fmt.Println(body.Password)

	err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(body.Password))
	if err != nil {

		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "incorrect  password!",
		})
	}

	fmt.Println(usr.ID)

	// create tokem

	// set to cookie

	// c.Cookie(&fiber.Cookie{
	// 	Name:     "adminauth",
	// 	Value:    tokenString,
	// 	HTTPOnly: true,
	// })
	tokenString, errMessage := utils.GenJwtToken("admin", usr.ID, 86400)
	if errMessage != "" {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": errMessage,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"access_tokem": tokenString,
	})

	// c.Send([]byte("cookie send"))
	// return c.Status(http.StatusOK).SendString("login success!")
}

func Validate(c *fiber.Ctx) error {
	fmt.Println("hhhh")

	user := c.Locals("id")
	fmt.Println(user)

	// use := c.Locals("id")
	// fmt.Println(use.(model.Admin).Name)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "login success",
	})

}
func UserManagement(c *fiber.Ctx) error {
	db := database.OpenDb()
	defer database.CloseDb(db)

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
func ViewUsers(c *fiber.Ctx) error {

	db := database.OpenDb()
	defer database.CloseDb(db)
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
	fmt.Println(address)

	infoWithoutAddress := utils.ExtractPersonalInfo(user)
	infowithPersonal := utils.ExtractAdresses(address)

	comb := utils.Combined(infoWithoutAddress, infowithPersonal)

	return c.Status(200).JSON(comb)

}

func Logout(c *fiber.Ctx) error {
	userId := c.Locals("id")
	usr_id := fmt.Sprintf("%v", userId)

	u64, err := strconv.ParseUint(usr_id, 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	tokenString, errMessage := utils.GenJwtToken("admin", uint(u64), 1)
	if errMessage != "" {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": errMessage,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"access_tokem": tokenString,
	})

}

func ViewOrders(c *fiber.Ctx) error {

	db := database.OpenDb()
	defer database.CloseDb(db)

	var orders []model.Order

	db.Find(&orders)
	fmt.Println(orders)

	extracttedOrderInfo := utils.ExtractOrderInfo(orders)

	return c.Status(200).JSON(extracttedOrderInfo)

}

func DeliveryStatusUpdate(c *fiber.Ctx) error {
	db := database.OpenDb()
	defer database.CloseDb(db)
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

		fmt.Println(order)
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