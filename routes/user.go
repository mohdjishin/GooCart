package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohdjishin/GoCart/controller"
	"github.com/mohdjishin/GoCart/middleware"
)

func UserRoute(app *fiber.App) {

	app.Post("/user_registration", controller.UserSignup) //json
	app.Post("/user_signin", controller.UserLogin)        //json
	app.Get("/home", middleware.RequreUserAuth, controller.Home)

	app.Post("/verification", middleware.RequreUserAuth, controller.Verification)      //json
	app.Post("/user/user_account", middleware.RequreUserAuth, controller.EditUserInfo) //json

}
