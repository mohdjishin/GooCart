package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohdjishin/GoCart/controller"
	"github.com/mohdjishin/GoCart/middleware"
)

func AdminRoute(app *fiber.App) {

	// app.Post("/admin/signup", controller.Signup)

	app.Post("/admin/login", controller.Login)                                                            //json
	app.Get("/admin_panel", middleware.RequireAdminAuth, controller.Validate)                             //nothing much
	app.Post("/admin_panel/add_product", middleware.RequireAdminAuth, controller.AddProducts)             //formdata
	app.Put("/admin_panel/products/edit_products/:id", middleware.RequireAdminAuth, controller.UpdatePro) // formdata
	app.Delete("admin_panel/products/delete_products/:id", middleware.RequireAdminAuth, controller.DelProduct)
	app.Get("admin_panel/view_users", middleware.RequireAdminAuth, controller.ViewUsers)

	app.Post("/admin_panel/user_management", middleware.RequireAdminAuth, controller.UserManagement)

	app.Get("/admin_panel/orders", middleware.RequireAdminAuth, controller.ViewOrders)
	app.Get("/admin_panel/logout", middleware.RequireAdminAuth, controller.Logout)

	app.Post("/admin_panel/delivery_status", middleware.RequreUserAuth, controller.DeliveryStatusUpdate)

	app.Post("/admin_panel/blockuser", middleware.RequireAdminAuth, controller.ManageUser)
	app.Post("/admin/refresh", controller.AdminRefresh)

}
