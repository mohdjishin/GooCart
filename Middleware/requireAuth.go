package Middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mohdjishin/fiberRESTApi/Database"
	"github.com/mohdjishin/fiberRESTApi/model"
)

func RequireAdminAuth(c *fiber.Ctx) error {
	db := Database.OpenDb()
	defer Database.CloseDb(db)

	// to get headers
	tkn := c.GetReqHeaders()

	tokenString := tkn["Authorization"]
	if tokenString == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "no token found",
		})
	}
	tokenString = tokenString[7:]

	// get the cookie
	// tokenString := c.Cookies("adminauth")
	// if tokenString == "" {
	// 	fmt.Println("failed")
	// 	return c.SendStatus(http.StatusUnauthorized)

	// }
	fmt.Println(tokenString)

	// decode/validate it

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header)
		}

		return []byte(os.Getenv("SECRET")), nil

	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check exp
		if claims["role"] != "admin" {

			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid token",
			})
		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return c.SendStatus(http.StatusUnauthorized)
		}

		// find the user with token sub

		user := new(model.Admin)

		db.First(&user, claims["sub"])

		if user.ID == 0 {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "user not found",
			})

		}

		// attach to req
		c.Locals("id", user.ID)

		// continue
		// fmt.Println(c.Locals("id"))

		c.Next()

	} else {
		fmt.Println(err)
		return c.SendStatus(http.StatusUnauthorized)
	}

	return nil
}
func RequreUserAuth(c *fiber.Ctx) error {
	db := Database.OpenDb()
	defer Database.CloseDb(db)

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
		// check exp

		if claims["role"] != "user" {

			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "no user privileges",
			})
		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return c.SendStatus(http.StatusUnauthorized)
		}

		// find the user with token sub

		user := new(model.Users)

		db.First(&user, claims["sub"])

		if user.ID == 0 {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "user not found",
			})

		}

		// attach to req
		c.Locals("id", user.ID)

		// continue
		// fmt.Println(c.Locals("id"))

		c.Next()

	} else {
		fmt.Println(err)
		return c.SendStatus(http.StatusUnauthorized)
	}

	return nil
}
