package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohdjishin/GoCart/controller"
	"github.com/mohdjishin/GoCart/middleware"
)

func AdminRoute(app *fiber.App) {

	// app.Post("/admin/signup", controller.Signup)

	admin := app.Group("/admin", middleware.RequireAdminAuth)

	app.Post("/admin/login", controller.Login) //json
	app.Post("/admin/refresh", controller.AdminRefresh)
	admin.Get("/admin/admin_panel", controller.Validate)                             //nothing much
	admin.Post("/admin/admin_panel/add_product", controller.AddProducts)             //formdata
	admin.Put("/admin/admin_panel/products/edit_products/:id", controller.UpdatePro) // formdata
	admin.Delete("/admin/admin_panel/products/delete_products/:id", controller.DelProduct)
	admin.Get("/adminadmin_panel/view_users", controller.ViewUsers)

	admin.Post("/admin/admin_panel/user_management", controller.UserManagement)

	admin.Get("/admin/admin_panel/orders", controller.ViewOrders)
	admin.Get("/admin/admin_panel/logout", controller.Logout)

	admin.Post("/admin/admin_panel/delivery_status", controller.DeliveryStatusUpdate)

	admin.Post("/admin/admin_panel/blockuser", controller.ManageUser)

}
