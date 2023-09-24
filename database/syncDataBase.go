package database

import (
	"log"

	"github.com/mohdjishin/GoCart/model"
)

var DB = NewDatabaseConnection()

func SyncDatabase() {
	db := DB.OpenDb()
	defer DB.CloseDb(db)
	err := db.AutoMigrate(&model.Admin{})
	errHandler(err)
	err = db.AutoMigrate(&model.Products{})
	errHandler(err)
	err = db.AutoMigrate(&model.ProductImage{})
	errHandler(err)
	err = db.AutoMigrate(&model.Users{})
	errHandler(err)
	err = db.AutoMigrate(&model.Address{})
	errHandler(err)
	err = db.AutoMigrate(&model.Cart{})
	errHandler(err)
	err = db.AutoMigrate(&model.Order{})
	errHandler(err)
	err = db.AutoMigrate(&model.CartTotal{})
	errHandler(err)

}

func errHandler(e error) {
	if e != nil {
		log.Println(e)
	}
}
