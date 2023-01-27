package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohdjishin/GoCart/controller"
	"github.com/mohdjishin/GoCart/middleware"
)

func AdminRoute(app *fiber.App) {

	// app.Post("/admin/signup", controller.Signup)

	app.Post("/admins_login", controller.Login) //json
	admin := app.Group("/admin", middleware.RequireAdminAuth)

	app.Post("/admins/refresh", controller.AdminRefresh)
	admin.Get("/admin_panel", controller.Validate)
	admin.Post("/admin_panel/add_product", controller.AddProducts)
	admin.Put("/admin_panel/products/edit_products/:id", controller.UpdatePro)
	admin.Delete("/admin_panel/products/delete_products/:id", controller.DelProduct)
	admin.Get("/admin_panel/view_users", controller.ViewUsers)

	admin.Post("admin_panel/user_management", controller.UserManagement)

	admin.Get("admin_panel/orders", controller.ViewOrders)
	admin.Get("admin_panel/logout", controller.Logout)

	admin.Post("admin_panel/delivery_status", controller.DeliveryStatusUpdate)

	admin.Post("/admin_panel/blockuser", controller.ManageUser)

}
