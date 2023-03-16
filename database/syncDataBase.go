package database

import "github.com/mohdjishin/GoCart/model"

var DB = NewDatabaseConnection()

func SyncDatabase() {
	db := DB.OpenDb()
	defer DB.CloseDb(db)
	db.AutoMigrate(&model.Admin{})

	db.AutoMigrate(&model.Products{})
	db.AutoMigrate(&model.ProductImage{})

	db.AutoMigrate(&model.Users{})
	db.AutoMigrate(&model.Address{})

	db.AutoMigrate(&model.Cart{})
	db.AutoMigrate(&model.Order{})
	db.AutoMigrate(&model.CartTotal{})
	// db.AutoMigrate(&model.CheckOut{})
	// db.AutoMigrate(&model.DirectIdCheckout{})
}
