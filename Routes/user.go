package Routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohdjishin/GoCart/Controller"
	"github.com/mohdjishin/GoCart/Middleware"
)

func UserRoute(app *fiber.App) {

	app.Post("/user_registration", Controller.UserSignup)
	app.Post("/user_signin", Controller.UserLogin)
	app.Get("/home", Middleware.RequreUserAuth, Controller.Home)

	app.Post("/verification", Middleware.RequreUserAuth, Controller.Verification)

}
