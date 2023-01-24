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

	Total           float64
	ShippmentStatus string `json:"status"`

	Aprovel bool `json:"approvel"`
}

type OrderRespAdmin struct {
	OrderID       uint   `json:"order_id"`
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"product_id"`
	Quantity      int    `json:"quantity"`
	Price         int    `json:"price"`
	Status        string `json:"status"`
	PaymentStatus bool   `json:"payment_status"`
}

type Invoice struct {
	OrderId     string
	ProductName string

	Quantity string
	Price    string
	Total    string
	Name     string
	Phone    string
}
