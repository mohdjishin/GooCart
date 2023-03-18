package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohdjishin/GoCart/controller"
	"github.com/mohdjishin/GoCart/middleware"
)

var p = controller.NewProduct()
var a = controller.NewAdmin()

func AdminRoute(app *fiber.App) {

	// app.Post("/admin/signup", a.Signup)

	app.Post("/admins_login", a.Login) //json
	admin := app.Group("/admin", middleware.RequireAdminAuth)

	app.Post("/admins/refresh", a.AdminRefresh)
	admin.Get("/admin_panel", a.Validate)
	admin.Post("/admin_panel/add_product", p.AddProducts)
	admin.Put("/admin_panel/products/edit_products/:id", p.UpdatePro)
	admin.Delete("/admin_panel/products/delete_products/:id", p.DelProduct)
	admin.Get("/admin_panel/view_users", a.ViewUsers)

	admin.Post("admin_panel/user_management", a.UserManagement)

	admin.Get("admin_panel/orders", a.ViewOrders)
	admin.Get("admin_panel/logout", a.Logout)

	admin.Post("admin_panel/delivery_status", a.DeliveryStatusUpdate)

	admin.Post("/admin_panel/blockuser", a.ManageUser)

}
