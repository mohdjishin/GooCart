package interfaces

import "github.com/gofiber/fiber/v2"

type IUser interface {
	UserSignup(*fiber.Ctx) error
	UserLogin(*fiber.Ctx) error
	Home(*fiber.Ctx) error
	Verification(*fiber.Ctx) error
	EditUserInfo(*fiber.Ctx) error
	AddToCart(*fiber.Ctx) error
	OrderFromCart(c *fiber.Ctx) error
	Checkout(*fiber.Ctx) error
	UserLogout(*fiber.Ctx) error
	Refresh(*fiber.Ctx) error
	GenerateInvoice(*fiber.Ctx) error
	RemoveFromCart(*fiber.Ctx) error
	First(*fiber.Ctx) error
}
