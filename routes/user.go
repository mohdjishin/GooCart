package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohdjishin/GoCart/controller"
	"github.com/mohdjishin/GoCart/middleware"
)

func UserRoute(app *fiber.App) {

	app.Get("/", controller.First)
	app.Post("/user/registration", controller.UserSignup) //json
	app.Post("/user/login", controller.UserLogin)         //json
	app.Post("/user/refresh", controller.Refresh)

	user := app.Group("/user", middleware.RequreUserAuth)

	user.Get("/home", controller.Home)

	user.Post("/verification", controller.Verification) //json
	user.Put("/user_account", controller.EditUserInfo)  //json
	user.Get("/view_products", p.ViewProducts)

	user.Get("/addtocart/:id", controller.AddToCart)
	user.Get("/removefromcart/:id", controller.RemoveFromCart)

	user.Post("/get_by_category", p.GetbyCategory)
	app.Get("/search/:key", p.SearchProduct)

	user.Get("/order_from_cart", controller.OrderFromCart)
	user.Get("/logout", controller.UserLogout)

	user.Get("/checkout", controller.Checkout)

	user.Get("/generate_invoice/:order_id", controller.GenerateInvoice)

}
