package database

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct{}

func NewDatabaseConnection() IDatabase {
	return Database{}
}

// var DB *gorm.DB

// func ConnectToDb() {
// 	var err error
// 	// import "gorm.io/driver/postgres"
// 	// ref: https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL

// 	DB, err = gorm.Open(postgres.Open(os.Getenv("DNS")), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println("err in db connection", err.Error())
// 	}

// }

func (Database) CloseDb(db *gorm.DB) {
	dbInstance, _ := db.DB()
	_ = dbInstance.Close()
}

func (Database) OpenDb() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DNS")), &gorm.Config{})
	if err != nil {

		log.Fatal("error in connecting database : ", err)
	}
	return db
}
