package database

import "github.com/mohdjishin/GoCart/model"

func SyncDatabase() {
	db := OpenDb()
	defer CloseDb(db)
	db.AutoMigrate(&model.Admin{})

	db.AutoMigrate(&model.Products{})
	db.AutoMigrate(&model.ProductImage{})

	db.AutoMigrate(&model.Users{})
	db.AutoMigrate(&model.Address{})

	db.AutoMigrate(&model.Cart{})
	db.AutoMigrate(&model.Order{})
}
