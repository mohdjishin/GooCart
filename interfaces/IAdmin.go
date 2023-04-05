package interfaces

import "github.com/gofiber/fiber/v2"

type IAdmin interface {
	Signup(*fiber.Ctx) error
	Login(*fiber.Ctx) error
	Validate(*fiber.Ctx) error
	UserManagement(*fiber.Ctx) error
	ViewUsers(*fiber.Ctx) error
	Logout(*fiber.Ctx) error
	ViewOrders(*fiber.Ctx) error
	DeliveryStatusUpdate(*fiber.Ctx) error
	ManageUser(*fiber.Ctx) error
	AdminRefresh(*fiber.Ctx) error
}
