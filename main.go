package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mohdjishin/GoCart/Database"
	"github.com/mohdjishin/GoCart/Routes"
)

func init() {

	Database.SyncDatabase()
}

func main() {

	app := fiber.New()
	app.Static("/images", "./public/upload")
	Routes.AdminRoute(app)
	Routes.UserRoute(app)
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))

}
