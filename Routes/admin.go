package Routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohdjishin/GoCart/Controller"
	"github.com/mohdjishin/GoCart/Middleware"
)

func AdminRoute(app *fiber.App) {

	// app.Post("/admin/signup", controller.Signup)

	app.Post("/admin/login", Controller.Login)                                                            //json
	app.Get("/admin_panel", Middleware.RequireAdminAuth, Controller.Validate)                             //nothing much
	app.Post("/admin_panel/add_product", Middleware.RequireAdminAuth, Controller.AddProducts)             //formdata
	app.Put("/admin_panel/products/edit_products/:id", Middleware.RequireAdminAuth, Controller.UpdatePro) // formdata
	app.Delete("admin_panel/products/delete_products/:id", Middleware.RequireAdminAuth, Controller.DelProduct)
	app.Get("admin_panel/view_users", Middleware.RequireAdminAuth, Controller.ViewUsers)

	app.Post("/admin_panel/user_management", Middleware.RequireAdminAuth, Controller.UserManagement)
}
