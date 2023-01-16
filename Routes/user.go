package Routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohdjishin/GoCart/Controller"
	"github.com/mohdjishin/GoCart/Middleware"
)

func UserRoute(app *fiber.App) {

	app.Post("/user_registration", Controller.UserSignup) //json
	app.Post("/user_signin", Controller.UserLogin)        //json
	app.Get("/home", Middleware.RequreUserAuth, Controller.Home)

	app.Post("/verification", Middleware.RequreUserAuth, Controller.Verification)      //json
	app.Post("/user/user_account", Middleware.RequreUserAuth, Controller.EditUserInfo) //json

}
