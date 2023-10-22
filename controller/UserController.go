package controller

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	I "github.com/mohdjishin/GoCart/interfaces"
	"github.com/mohdjishin/GoCart/model"
	utils "github.com/mohdjishin/GoCart/utils"
	"golang.org/x/crypto/bcrypt"
)

var BillGen = utils.NewBillGenerator()
var Token = utils.NewToken()

type User struct{}

func NewUserFunc() I.IUser {
	return &User{}
}

func (*User) UserSignup(c *fiber.Ctx) error {

	db := DB.OpenDb()
	defer DB.CloseDb(db)

	user := new(model.Users)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	if err := utils.ValidateStruct(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	address := new(model.Address)
	if err := utils.ValidateStruct(address); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := c.BodyParser(address); err != nil {
		return c.Status(500).SendString(err.Error())
	}

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

	res := db.Save(&user)

	if res.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create account",
		})

	}

	if true {
		stst := utils.SendOtp("+" + user.CountryCode + user.Phone)
		if stst {
			fmt.Println("otp sented sccessfully")
		}

	} else {
		err = utils.SendSMSOTP(("+" + user.CountryCode + user.Phone), OTP)
		if err != nil {
			fmt.Println(err)

		}

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

func (*User) UserLogin(c *fiber.Ctx) error {
	db := DB.OpenDb()
	defer DB.CloseDb(db)

	body := new(model.Users)

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
	if usr.Blocked == true {

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "please contact admin",
		})

	}
	fmt.Println(usr.ID)

	uuidv4, _ := uuid.NewRandom()

	err = db.Model(&usr).Where("id = ?", usr.ID).Update("refresh", uuidv4.String()).Error
	if err != nil {
		fmt.Println(err)

	}

	tokenString, errMessage := Token.GenJwtToken("user", usr.ID, 86400)
	if errMessage != "" {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": errMessage,
		})
	}
	fmt.Println(tokenString)
	if err != nil {

		_ = utils.InternalServerError("Issue generating token", c)
	}

	// set to cookie

	// c.Cookie(&fiber.Cookie{
	// 	Name:     "userauth",
	// 	Value:    tokenString,
	// 	HTTPOnly: true,
	// })

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"access_tokem":  tokenString,
		"refresh_token": uuidv4,
	})

}

func (*User) Home(c *fiber.Ctx) error {

	user := c.Locals("id")
	fmt.Println(user)
	fmt.Println("Helo")

	return nil
}

func (*User) Verification(c *fiber.Ctx) error {

	db := DB.OpenDb()

	defer DB.CloseDb(db)
	status := false

	userId := c.Locals("id")

	otp := new(model.OTP)

	if err := c.BodyParser(otp); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	user := new(model.Users)

	db.First(&user, userId)

	if true {
		status := utils.CheckOtp(("+" + user.CountryCode + user.Phone), otp.OTP)
		user.Verified = status

	} else {
		if user.OTP == otp.OTP {
			status = true
			err := utils.WelcomeMsg("+" + user.CountryCode + user.Phone)

			if err != nil {
				fmt.Println(err)
			}
		}

		user.Verified = status

	}

	db.Save(&user)

	if !status {
		return c.Status(200).JSON(fiber.Map{"message": "otp verification failed"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "phone number verified"})

}

func (*User) EditUserInfo(c *fiber.Ctx) error {
	db := DB.OpenDb()
	defer DB.CloseDb(db)
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

	userInfo := new(model.Users)

	db.Find(&userInfo, userId)

	addressInfo := new(model.Address)

	db.Find(&addressInfo, userId)

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

func (*User) AddToCart(c *fiber.Ctx) error {
	db := DB.OpenDb()
	defer DB.CloseDb(db)
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

		cart.ProductId = uint(prod.ID)
		cart.Price = prod.Price
		cart.Image = prodImg.ImageOne
		cart.Quantity = 1
		cart.Total = (cart.Price * float64(cart.Quantity))
		res := db.Create(&cart)
		if res.Error != nil {
			_ = utils.InternalServerError("failed adding to cart", c)

		}
		return c.Status(200).JSON(cart)
	}
	cart.CartID = i
	cart.UserId = usr_id

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

		_ = utils.InternalServerError("failed adding to cart", c)

	}

	return c.Status(200).JSON(cart)
}

func (*User) OrderFromCart(c *fiber.Ctx) error {
	db := DB.OpenDb()
	defer DB.CloseDb(db)
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
		orders.ShippmentStatus = "processing"
		userID, err := strconv.ParseUint(c.UserId, 10, 32)
		if err != nil {
			fmt.Println(err)
		}

		orders.UserID = uint(userID)
		orders.AddressID = uint(userID)
		orders.Price = c.Price
		orders.ShippmentStatus = "processing"

		orders.Quantity = c.Quantity

		orders.Total = c.Total * float64(c.Quantity)
		db.Create(&orders)
		db.Delete(&c)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "ordered",
	})
}

func (*User) Checkout(c *fiber.Ctx) error {
	db := DB.OpenDb()
	defer DB.CloseDb(db)
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

	paymentLink := utils.Payment("products", total)
	return c.JSON(fiber.Map{
		"cart":         cart,
		"address":      address,
		"total":        total,
		"payment_link": paymentLink,
	})
}

func (*User) UserLogout(c *fiber.Ctx) error {

	userId := c.Locals("id")
	usr_id := fmt.Sprintf("%v", userId)

	u64, err := strconv.ParseUint(usr_id, 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	tokenString, errMessage := Token.GenJwtToken("user", uint(u64), 1)
	if errMessage != "" {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": errMessage,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"access_token": tokenString,
	})

}

func (*User) Refresh(c *fiber.Ctx) error {
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

	err, accessToken, rfToken := Token.RefreshToken(db, rt.Refresh_token, rt.Access_token)

	if err != "" {
		return c.Status(400).JSON(fiber.Map{"message": err})
	}
	return c.Status(200).JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": rfToken,
	})
}

func (*User) GenerateInvoice(c *fiber.Ctx) error {

	db := DB.OpenDb()
	defer DB.CloseDb(db)
	orderId := c.Params("order_id")

	bill := new(model.Invoice)
	user := new(model.Users)

	var order model.Order

	var prod model.Products

	db.Find(&order, "order_id =?", orderId)

	db.First(&user, order.UserID)

	db.Find(&prod, order.ProductID)

	bill.Name = user.Name
	bill.Phone = user.CountryCode + user.Phone
	bill.OrderId = order.OrderId

	fmt.Println(order)
	total := prod.Price * float64(order.Quantity)

	bill.ProductName = prod.Product_Name

	bill.Price = fmt.Sprintf("%v", order.Price)
	bill.Quantity = fmt.Sprintf("%v", order.Quantity)
	bill.Total = fmt.Sprintf("%v", total)

	filen := BillGen.GenerateInvoice(*bill)

	st, url := utils.UploadPDFToS3[string]("/home/mohdjishin/brototype/GoCart/media/pdf/"+filen, filen)
	if st {
		fmt.Println("uploaded")
		err := os.Remove("/home/mohdjishin/brototype/GoCart/media/pdf/" + filen)

		if err != nil {
			fmt.Println(err)
		}
	}
	// combineOrderAndProd := utils.CombinedPRoductOrder(order, prod)

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"invoice": url,
	})
}

func (*User) RemoveFromCart(c *fiber.Ctx) error {

	db := DB.OpenDb()
	defer DB.CloseDb(db)
	user_Id := c.Locals("id")
	usr_id := fmt.Sprintf("%v", user_Id)

	prod := new(model.Products)

	product_Id := c.Params("id")
	db.First(&prod, product_Id)

	// db.Find(&prod, proId)

	cart := new(model.Cart)

	prodImg := new(model.ProductImage)

	cartTotal := new(model.CartTotal)

	res := db.First(&prodImg, "product_id =? ", product_Id)
	if res.Error != nil {
		return c.Status(200).JSON(fiber.Map{

			"message": "no product found with pro_id :" + product_Id,
		})

	}

	res = db.Where("user_id = ? AND product_id >= ?", usr_id, product_Id).First(&cart)

	if res.Error != nil {
		fmt.Println("ok, no product found..")
	}

	db.Where("cart_id = ?", cart.CartID).First(&cart)

	if cart.ID == 0 {
		fmt.Println("no item found")
	}

	if cart.Quantity > 1 {

		db.Model(&cart).Update("quantity", cart.Quantity-1)
		ttl := cart.Quantity * int(cart.Price)
		db.Model(&cart).Update("total", ttl)
		db.Model(&cartTotal).Update("total", ttl)

		return c.Status(200).JSON(fiber.Map{
			"message": "item deteled",
			"cart":    cart,
		})

	} else if cart.Quantity == 1 {
		db.Delete(&cart)
		db.Delete(&cartTotal)
		return c.Status(200).JSON(fiber.Map{
			"message": "item removed",
			"cart":    cart,
		})

	}
	fmt.Println(cartTotal.CartID)
	db.Delete(&cartTotal)
	return c.Status(200).JSON(fiber.Map{
		"message": "no item found",
	})
}

func (*User) First(c *fiber.Ctx) error {
	return c.Status(200).SendString(" please visit - https://github.com/mohdjishin/GooCart")
}
