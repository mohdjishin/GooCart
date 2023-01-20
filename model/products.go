package model

import (
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Product_Category string  `json:"pro_category"`
	Product_Name     string  `json:"pro_name"`
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
	CartID    int `gorm:"auto_increment"`
	UserId    string
	ProductId uint
	Quantity  int
	Image     string
	Price     float64
	Total     float64
}

type ProductInfo struct {
	ID              int
	ProductCategory string
	ProductName     string
	Price           float64
}

type ProductImageInfo struct {
	ProductId  int
	ImageOne   string
	ImageTwo   string
	ImageThree string
}

type CombinedProductInfo struct {
	ID              int     `json:"id"`
	ProductCategory string  `json:"pro_category"`
	ProductName     string  `json:"pro_name"`
	Price           float64 `json:"price"`
	ImageOne        string  `json:"img_one"`
	ImageTwo        string  `json:"img_two"`
	ImageThree      string  `json:"img_three"`
}

type CartTotal struct {
	gorm.Model
	CartID int     `json:"-"`
	Total  float64 `json:"-"`
}
