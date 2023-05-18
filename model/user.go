package model

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model

	Name        string `json:"name" validate:"required"`
	Username    string `json:"username"  gorm:"index;unique" validate:"required"`
	Email       string `json:"email" validate:"required,email" `
	Password    string `json:"password" validate:"required" `
	Phone       string `json:"phone" validate:"required"`
	CountryCode string `json:"country_code" validate:"required"`
	Verified    bool
	OTP         string
	Status      bool   `json:"status"`
	Blocked     bool   `json:"block_status"`
	Refresh     string `json:"refresh"`
}

type Address struct {
	gorm.Model
	UserId uint `gorm:"index;unique" `
	// Users     Users
	HouseName string `json:"house_name" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
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

type PersonalInformation struct {
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
