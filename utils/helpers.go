package utils

import (
	"fmt"
	"regexp"

	"github.com/mohdjishin/GoCart/model"
)

type information struct {
	UserId       int
	Name         string
	Username     string
	Phone        string
	Mail         string
	Verification bool
}

type Combine struct {
	UserId       int
	Name         string
	Username     string
	Phone        string
	Mail         string
	Verification bool
	Housename    string
	Street       string
	City         string
	Pin          string
	State        string
}

type Extractaddress struct {
	UserID    uint
	Housename string
	Street    string
	City      string
	Pin       string
	State     string
}

func ExtractPersonalInfo(users []model.Users) []information {
	var newUsers []information
	for _, user := range users {
		newUsers = append(newUsers, information{
			UserId:       int(user.ID),
			Name:         user.Name,
			Username:     user.Username,
			Phone:        user.Phone,
			Mail:         user.Email,
			Verification: user.Verified,
		})
	}
	return newUsers
}
func ExtractAdresses(addr []model.Address) []Extractaddress {
	var newUsers []Extractaddress
	for _, addr := range addr {
		newUsers = append(newUsers, Extractaddress{
			UserID:    addr.UserId,
			Housename: addr.HouseName,
			Street:    addr.Street,
			City:      addr.City,
			Pin:       addr.Pin,
			State:     addr.State,
		})
	}
	return newUsers
}

func Combined(users []information, addresses []Extractaddress) []Combine {
	var combined []Combine
	for _, u := range users {
		for _, a := range addresses {

			if u.UserId == int(a.UserID) {
				fmt.Println("he;p")
				combined = append(combined, Combine{UserId: u.UserId, Username: u.Username, Name: u.Name, Mail: u.Mail, Phone: u.Phone, Verification: u.Verification, Housename: a.Housename, Street: a.Street, Pin: a.Pin, City: a.City, State: a.State})
			}
		}
	}
	fmt.Println(combined)
	return combined
}

func CheckComplexityOFPassword(password string) bool {
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString
	hasLower := regexp.MustCompile(`[a-z]`).MatchString
	hasSymbol := regexp.MustCompile(`[!@#\$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString
	return hasNumber(password) && hasUpper(password) && hasLower(password) && hasSymbol(password)
}
