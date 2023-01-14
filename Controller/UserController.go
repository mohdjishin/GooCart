package Controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mohdjishin/fiberRESTApi/Database"
	utils "github.com/mohdjishin/fiberRESTApi/Utils"
	"github.com/mohdjishin/fiberRESTApi/model"
	"golang.org/x/crypto/bcrypt"
)

func UserSignup(c *fiber.Ctx) error {

	db := Database.OpenDb()
	defer Database.CloseDb(db)
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
	utils.SendOtp(("+" + user.CountryCode + user.Phone))

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
	db := Database.OpenDb()
	defer Database.CloseDb(db)

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
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{

			"message": "Issue generating token",
		})
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

	db := Database.OpenDb()

	defer Database.CloseDb(db)

	userId := c.Locals("id")

	otp := new(model.OTP)

	if err := c.BodyParser(otp); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	user := new(model.Users)

	db.First(&user, userId)

	status := utils.CheckOtp(("+" + user.CountryCode + user.Phone), otp.OTP)

	user.Verified = status

	db.Save(&user)

	if !status {
		return c.Status(200).JSON(fiber.Map{"message": "otp verification failed"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "phone number verified"})

}
