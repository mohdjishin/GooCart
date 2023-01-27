package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mohdjishin/GoCart/database"
	"github.com/mohdjishin/GoCart/routes"
)

func init() {

	database.SyncDatabase()

}

func main() {

	app := fiber.New()

	routes.AdminRoute(app)

	routes.UserRoute(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))

}
