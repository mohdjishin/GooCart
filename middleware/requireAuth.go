package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	Database "github.com/mohdjishin/GoCart/database"
	"github.com/mohdjishin/GoCart/model"
)

var DB = Database.NewDatabaseConnection()

func RequireAdminAuth(c *fiber.Ctx) error {
	db := DB.OpenDb()
	defer DB.CloseDb(db)

	tkn := c.GetReqHeaders()

	tokenString := tkn["Authorization"]
	if tokenString == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "no token found",
		})
	}
	tokenString = tokenString[7:]

	fmt.Println(tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header)
		}

		return []byte(os.Getenv("SECRET")), nil

	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if claims["role"] != "admin" {

			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid token",
			})
		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return c.SendStatus(http.StatusUnauthorized)
		}

		user := new(model.Admin)

		db.First(&user, claims["sub"])

		if user.ID == 0 {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "user not found",
			})

		}

		c.Locals("id", user.ID)

		err = c.Next()
		if err != nil {
			log.Println(err)
		}
	} else {
		fmt.Println(err)
		return c.SendStatus(http.StatusUnauthorized)
	}

	return nil
}
func RequreUserAuth(c *fiber.Ctx) error {
	db := DB.OpenDb()
	defer DB.CloseDb(db)

	fmt.Println("In middleware")

	headers := c.GetReqHeaders()

	tokenString := headers["Authorization"]
	if tokenString == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "no token found",
		})
	}
	tokenString = tokenString[7:]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header)
		}

		return []byte(os.Getenv("SECRET")), nil

	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if claims["role"] != "user" {

			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "no user privileges",
			})
		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return c.SendStatus(http.StatusUnauthorized)
		}

		user := new(model.Users)

		db.First(&user, claims["sub"])

		if user.ID == 0 {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "user not found",
			})

		}

		c.Locals("id", user.ID)

		err = c.Next()
		if err != nil {
			log.Println(err)
		}

	} else {
		fmt.Println(err)
		return c.SendStatus(http.StatusUnauthorized)
	}

	return nil
}
