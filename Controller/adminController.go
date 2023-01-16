package Controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mohdjishin/GoCart/Database"
	"github.com/mohdjishin/GoCart/model"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {
	db := Database.OpenDb()

	Database.CloseDb(db)
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

	res := db.Save(&usr)

	if res.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create account",
		})

	}
	return c.Status(200).SendString("account created")
}

func Login(c *fiber.Ctx) error {
	db := Database.OpenDb()
	defer Database.CloseDb(db)
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

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "incorrect  password!",
		})
	}

	fmt.Println(usr.ID)

	// create tokem
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{

		"role": "admin",
		"sub":  usr.ID,
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
	// 	Name:     "adminauth",
	// 	Value:    tokenString,
	// 	HTTPOnly: true,
	// })

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
	db := Database.OpenDb()
	defer Database.CloseDb(db)

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
