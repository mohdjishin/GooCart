package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mohdjishin/fiberRESTApi/Database"
	"github.com/mohdjishin/fiberRESTApi/Routes"
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
