package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mohdjishin/GoCart/database"
	"github.com/mohdjishin/GoCart/routes"
	"github.com/mohdjishin/GoCart/utils"
)

func init() {
	time.Sleep(time.Second * 7)
	database.SyncDatabase()

}

func main() {

	app := fiber.New()

	routes.AdminRoute(app)

	routes.UserRoute(app)

	utils.ListenAndShutdown(app)
}
