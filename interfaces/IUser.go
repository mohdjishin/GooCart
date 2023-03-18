package interfaces

import "github.com/gofiber/fiber/v2"

type IUser interface {
	UserSignup(*fiber.Ctx) error
	UserLogin(c *fiber.Ctx) error
	Home(c *fiber.Ctx) error
	Verification(c *fiber.Ctx) error
	EditUserInfo(c *fiber.Ctx) error
	AddToCart(c *fiber.Ctx) error
	OrderFromCart(c *fiber.Ctx) error
	Checkout(c *fiber.Ctx) error
	UserLogout(c *fiber.Ctx) error
	Refresh(c *fiber.Ctx) error
	GenerateInvoice(c *fiber.Ctx) error
	RemoveFromCart(c *fiber.Ctx) error
	First(c *fiber.Ctx) error
}
