package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenJwtToken(role string, userId uint, duration int) (string, string) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{

		"role": role,
		"sub":  userId,
		"exp":  time.Now().Add(time.Second * time.Duration(duration)).Unix(),
	})

	fmt.Println(token.Valid)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	fmt.Println(tokenString)
	if err != nil {
		return "", "Issue generating token"
	}

	return tokenString, ""

}
