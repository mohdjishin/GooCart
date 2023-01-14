package model

import (
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Product_Category string  `json:"product_category"`
	Product_Name     string  `json:"product_name"`
	Price            float64 `json:"price"`
}

type ProductImage struct {
	gorm.Model
	ProductId uint   `json:"product_id"`
	ImageOne  string `json:"img_one"`
	ImgTwo    string `json:"img_two"`
	ImgThree  string `json:"img_three"`
}

type Cart struct {
	gorm.Model
	UserId    uint
	ProductId uint
	Quantity  int
}
