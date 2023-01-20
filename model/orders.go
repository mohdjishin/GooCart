package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	OrderId string `json:"order_id" gorm:"primary_key;unique;not null"`
	UserID  uint   `gorm:"not null" json:"user_id"`

	ProductID uint `json:"product_id"`

	AddressID uint `json:"address_id"`

	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`

	Total  float64
	Status string `json:"status"`

	PaymentStatus string
	Aproveled     bool
}

type InstentOrder struct {
}
