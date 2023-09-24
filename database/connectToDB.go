package database

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	I "github.com/mohdjishin/GoCart/interfaces"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct{}

func NewDatabaseConnection() I.IDatabase {
	return Database{}
}

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
