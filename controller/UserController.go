package controller

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/mohdjishin/GoCart/database"
	"github.com/mohdjishin/GoCart/model"
	utils "github.com/mohdjishin/GoCart/utils"
	"golang.org/x/crypto/bcrypt"
)

func UserSignup(c *fiber.Ctx) error {

	db := database.OpenDb()
	defer database.CloseDb(db)
	// get the username and password
	user := new(model.Users)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	address := new(model.Address)
	if err := c.BodyParser(address); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// fmt.Println(user.ID)
	// fmt.Println(user.Name)
	// fmt.Println(user.Username)
	// fmt.Println(user.Email)
	// fmt.Println(user.Password)
	// fmt.Println(user.CountryCode)
	// fmt.Println(user.Phone)

	fmt.Println("+" + user.CountryCode + user.Phone)

	// utils.SendOtp(("+" + user.CountryCode + user.Phone))

	OTP := strconv.FormatUint(uint64(uuid.New().ID()), 10)[:6]

	user.OTP = OTP

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "not able to hash pasword",
		})

	}
	fmt.Println(hash)

	user.Password = string(hash)

	// create the user

	res := db.Save(&user)

	if res.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create account",
		})

	}
	err = utils.SendSMSOTP(("+" + user.CountryCode + user.Phone), OTP)
	if err != nil {
		fmt.Println(err)

	}

	address.UserId = user.ID

	res = db.Save(&address)

	if res.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create account",
		})

	}
	return c.Status(200).SendString("account created")
}

func UserLogin(c *fiber.Ctx) error {
	db := database.OpenDb()
	defer database.CloseDb(db)

	body := new(model.Users)
	// take data from req
	if err := c.BodyParser(body); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	usr := new(model.Users)
	db.First(&usr, "username = ?", body.Username)
	if usr.ID == 0 {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid  usermame or password",
		})
	}

	// check pass

	err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(body.Password))
	if err != nil {

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "incorrect  password!",
		})
	}

	if usr.Status {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "you have been restricted",
		})
	}

	fmt.Println(usr.ID)

	// create tokem
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  usr.ID,
		"role": "user",
		"exp":  time.Now().Add(time.Hour * 5).Unix(),
	})

	fmt.Println(token.Valid)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	fmt.Println(tokenString)
	if err != nil {

		utils.InternalServerError("Issue generating token", c)
	}

	// set to cookie

	// c.Cookie(&fiber.Cookie{
	// 	Name:     "userauth",
	// 	Value:    tokenString,
	// 	HTTPOnly: true,
	// })

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"access_tokem": tokenString,
	})

	// c.Send([]byte("cookie send"))
	// return c.Status(http.StatusOK).SendString("login success!")
}

func Home(c *fiber.Ctx) error {

	user := c.Locals("id")
	fmt.Println(user)
	fmt.Println("Helo")

	return nil
}

func Verification(c *fiber.Ctx) error {

	db := database.OpenDb()

	defer database.CloseDb(db)
	status := false

	userId := c.Locals("id")

	otp := new(model.OTP)

	if err := c.BodyParser(otp); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	user := new(model.Users)

	db.First(&user, userId)

	// status := utils.CheckOtp(("+" + user.CountryCode + user.Phone), otp.OTP)
	if user.OTP == otp.OTP {
		status = true
		utils.WelcomeMsg("+" + user.CountryCode + user.Phone)
	}

	user.Verified = status

	db.Save(&user)

	if !status {
		return c.Status(200).JSON(fiber.Map{"message": "otp verification failed"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "phone number verified"})

}

func EditUserInfo(c *fiber.Ctx) error {
	db := database.OpenDb()
	defer database.CloseDb(db)
	userId := c.Locals("id")

	// get user info from req
	user := new(model.Users)
	address := new(model.Address)

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	// get address info from req
	if err := c.BodyParser(address); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	fmt.Println(user.Name)

	fmt.Println(user.Password)
	fmt.Println(user.Phone)

	fmt.Println(user.CountryCode)

	fmt.Println(address.HouseName)

	fmt.Println(address.Street)

	fmt.Println(address.City)
	fmt.Println(address.State)
	fmt.Println(address.Pin)

	userInfo := new(model.Users)

	db.Find(&userInfo, userId)
	fmt.Println(userInfo)

	addressInfo := new(model.Address)

	db.Find(&addressInfo, userId)

	fmt.Println(addressInfo)

	// get password and hash it

	if user.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "not able to hash pasword",
			})

		}
		userInfo.Password = string(hash)
	}

	if user.Phone != "" {
		userInfo.Phone = user.Phone
		userInfo.Verified = false
		userInfo.CountryCode = user.CountryCode

	}
	if user.Name != "" {
		userInfo.Name = user.Name
	}
	addressInfo.HouseName = address.HouseName
	addressInfo.Street = address.Street
	addressInfo.State = address.State
	addressInfo.City = address.City
	addressInfo.Pin = address.Pin

	db.Save(&userInfo)
	db.Save(&addressInfo)
	res := model.UserInfo{
		Name:        userInfo.Name,
		Username:    userInfo.Username,
		Email:       userInfo.Email,
		Password:    "Sensitive information",
		Phone:       userInfo.Phone,
		Verified:    userInfo.Verified,
		CountryCode: userInfo.CountryCode,
		HouseName:   addressInfo.HouseName,
		Street:      addressInfo.Street,
		City:        addressInfo.City,
		State:       addressInfo.State,
		Pin:         addressInfo.Pin,
	}

	return c.Status(200).JSON(res)
}

func AddToCart(c *fiber.Ctx) error {
	db := database.OpenDb()
	defer database.CloseDb(db)
	user_Id := c.Locals("id")
	usr_id := fmt.Sprintf("%v", user_Id)

	prod := new(model.Products)

	product_Id := c.Params("id")
	db.First(&prod, product_Id)
	// db.Find(&prod, proId)

	cart := new(model.Cart)

	prodImg := new(model.ProductImage)

	res := db.First(&prodImg, "product_id =? ", product_Id)
	if res.Error != nil {
		return c.Status(200).JSON(fiber.Map{

			"message": "no product found with pro_id :" + product_Id,
		})

	}

	res = db.Where("user_id = ? AND product_id >= ?", usr_id, product_Id).First(&cart)
	if res.Error != nil {
		fmt.Println("ok, new product adding...")
	}
	i, err := strconv.Atoi(usr_id)
	if err != nil {
		fmt.Println(err)
	}
	if cart.ID == 0 {

		cart.UserId = usr_id
		cart.CartID = i
		fmt.Println("hhhh")
		cart.ProductId = uint(prod.ID)
		cart.Price = prod.Price
		cart.Image = prodImg.ImageOne
		cart.Quantity = 1
		cart.Total = (cart.Price * float64(cart.Quantity))
		res := db.Create(&cart)
		if res.Error != nil {
			utils.InternalServerError("failed adding to cart", c)

		}
		return c.Status(200).JSON(cart)
	}
	cart.CartID = i
	cart.UserId = usr_id
	fmt.Println("hhhh")
	cart.ProductId = uint(prod.ID)

	cart.Image = prodImg.ImageOne
	if cart.Quantity == 0 {

		cart.Quantity = 1
	} else {
		cart.Quantity = cart.Quantity + 1
	}
	cart.Price = prod.Price

	cart.Total = (cart.Price * float64(cart.Quantity))
	// fmt.Println(prod.Product_Name)
	res = db.Save(&cart)
	if res.Error != nil {

		utils.InternalServerError("failed adding to cart", c)

	}

	return c.Status(200).JSON(cart)
}

func OrderFromCart(c *fiber.Ctx) error {
	db := database.OpenDb()
	defer database.CloseDb(db)
	user_Id := c.Locals("id")
	usr_id := fmt.Sprintf("%v", user_Id)

	// products := []model.Products{}

	cart := []model.Cart{}

	orders := new(model.Order)

	db.Where("user_id = ?", usr_id).Find(&cart)
	// fmt.Println(cart)

	for _, c := range cart {
		id := uuid.New().ID()
		str := strconv.Itoa(int(id))
		orders.OrderId = str[:7]
		orders.ProductID = c.ProductId
		userID, err := strconv.ParseUint(c.UserId, 10, 32)
		if err != nil {
			fmt.Println(err)
		}

		orders.UserID = uint(userID)
		orders.AddressID = uint(userID)
		orders.Price = c.Price

		orders.Quantity = c.Quantity
		orders.PaymentStatus = "pending"

		orders.Total = c.Total * float64(c.Quantity)
		db.Create(&orders)
		db.Delete(&c)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "ordered",
	})
}

func Checkout(c *fiber.Ctx) error {
	db := database.OpenDb()
	defer database.CloseDb(db)
	uid := c.Locals("id")
	user_id := fmt.Sprintf("%v", uid)
	cartTotal := new(model.CartTotal)

	var cart []model.Cart
	address := new(model.Address)

	res := db.Where("user_id = ? ", user_id).Find(&cart).Error
	if res != nil {
		fmt.Println(res)
	}
	total := 0.0
	for _, v := range cart {

		total = total + v.Total

	}
	i, err := strconv.Atoi(user_id)
	if err != nil {
		fmt.Println(err)
	}

	db.First(&cartTotal, "cart_id = ?", user_id)
	fmt.Println(cartTotal.ID)
	if cartTotal.ID == 0 {
		cartTotal.CartID = i
		cartTotal.Total = total

		db.Save(cartTotal)
	} else {

		db.Model(&model.CartTotal{}).Where("cart_id = ?", user_id).Update("total", total)
	}

	res = db.Where("user_id = ? ", user_id).Find(&address).Error
	fmt.Println(address)

	fmt.Println(total)
	fmt.Println(cart)
	return c.JSON(fiber.Map{
		"cart":    cart,
		"address": address,
		"total":   total,
	})
}
