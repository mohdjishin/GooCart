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

	app.Post("/verification", middleware.RequreUserAuth, controller.Verification)     //json
	app.Put("/user/user_account", middleware.RequreUserAuth, controller.EditUserInfo) //json
	app.Get("/user/view_products", middleware.RequreUserAuth, controller.ViewProducts)

	app.Get("/user/addtocart/:id", middleware.RequreUserAuth, controller.AddToCart)

	app.Post("/user/get_by_category", middleware.RequreUserAuth, controller.GetbyCategory)
	app.Get("/user/search/:key", controller.SearchProduct)

	app.Get("user/order_from_cart", middleware.RequreUserAuth, controller.OrderFromCart)
	// app.Get("/user/instant_buy_checkout/:id", middleware.RequreUserAuth, controller.BuytoCheckout)
	// app.Get("/user/remove_from_checkout/:id", middleware.RequreUserAuth, controller.RemovetoCheckout)

	app.Get("/user/checkout", middleware.RequreUserAuth, controller.Checkout)
}
