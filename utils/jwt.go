package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	I "github.com/mohdjishin/GoCart/interfaces"
	"github.com/mohdjishin/GoCart/model"
	"gorm.io/gorm"
)

type token struct{}

func NewToken() I.IToken {
	return &token{}
}

func (*token) GenJwtToken(role string, userId uint, duration int) (string, string) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{

		"role": role,
		"sub":  userId,
		"exp":  time.Now().Add(time.Second * time.Duration(duration)).Unix(),
	})

	fmt.Println(token.Valid)
	fmt.Println(token.Valid)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	fmt.Println(tokenString)
	if err != nil {
		return "", "Issue generating token"
	}

	return tokenString, ""

}

func (t *token) RefreshToken(db *gorm.DB, refresh string, accessToken string) (string, string, string) {

	tokenString := accessToken
	if tokenString == "" {
		return "no token found", "", ""
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header)
		}

		return []byte(os.Getenv("SECRET")), nil

	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if claims["role"] != "user" {

			return "no user privileges", "", ""

		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return "unauthorized", "", ""
		}

		user := new(model.Users)

		db.First(&user, claims["sub"])

		if user.ID == 0 {

			return "user not found", "", ""

		}

		if user.Refresh == refresh {
			newToken, err := t.GenJwtToken("user", user.ID, 86400)
			if err != "" {
				return err, "", ""
			}

			uuidv4, _ := uuid.NewRandom()

			errr := db.Model(&user).Where("id = ?", user.ID).Update("refresh", uuidv4).Error
			if errr != nil {
				fmt.Println(err)

			}

			return "", newToken, uuidv4.String()
		}

	}
	fmt.Println(err)
	return "unauthorized", "", ""

}

func (t *token) AdminRefreshToken(db *gorm.DB, refresh string, accessToken string) (string, string, string) {

	tokenString := accessToken
	if tokenString == "" {
		return "no token found", "", ""
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header)
		}

		return []byte(os.Getenv("SECRET")), nil

	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if claims["role"] != "admin" {

			return "no user privileges", "", ""

		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return "unauthorized", "", ""
		}

		user := new(model.Admin)

		db.First(&user, claims["sub"])

		if user.ID == 0 {

			return "user not found", "", ""

		}

		if user.Refresh == refresh {
			newToken, err := t.GenJwtToken("admin", user.ID, 86400)
			if err != "" {
				return err, "", ""
			}

			uuidv4, _ := uuid.NewRandom()

			errr := db.Model(&user).Where("id = ?", user.ID).Update("refresh", uuidv4).Error
			if errr != nil {
				fmt.Println(err)

			}

			return "", newToken, uuidv4.String()
		}

	}
	fmt.Println(err)
	return "unauthorized", "", ""

}
