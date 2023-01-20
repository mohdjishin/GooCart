package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model

	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
