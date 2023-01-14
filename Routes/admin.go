package Routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohdjishin/fiberRESTApi/Controller"
	"github.com/mohdjishin/fiberRESTApi/Middleware"
)

func AdminRoute(app *fiber.App) {

	// app.Post("/admin/signup", controller.Signup)

	app.Post("/admin/login", Controller.Login)
	app.Get("/admin_panel", Middleware.RequireAdminAuth, Controller.Validate)
	app.Post("/admin_panel/add_product", Middleware.RequireAdminAuth, Controller.AddProducts)
	app.Put("/admin_panel/products/edit_products/:id", Middleware.RequireAdminAuth, Controller.UpdatePro)
	app.Delete("admin_panel/products/delete_products/:id", Middleware.RequireAdminAuth, Controller.DelProduct)
	app.Get("admin_panel/view_users", Middleware.RequireAdminAuth, Controller.ViewUsers)
}
