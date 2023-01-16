package model

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model

	Name        string `json:"name"`
	Username    string `json:"username"  gorm:"index;unique"`
	Email       string `json:"email"  `
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	CountryCode string `json:"country_code"`
	Verified    bool
	Status      bool `json:"status"`
}

type Address struct {
	gorm.Model
	UserId uint `gorm:"index;unique"`
	// Users     Users
	HouseName string `json:"house_name"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	Pin       string `json:"pin"`
}

type OTP struct {
	OTP string `json:"otp"`
}

type UserInfo struct {
	Name        string
	Username    string
	Email       string
	Password    string
	Phone       string
	CountryCode string
	Verified    bool

	HouseName string
	Street    string
	City      string
	State     string
	Pin       string
}
