package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model

	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`

	Refresh string `json:"refresh"`
}
