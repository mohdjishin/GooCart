package controller

import "github.com/gofiber/fiber/v2"

type IAdmin interface {
	Signup(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Validate(c *fiber.Ctx) error
	UserManagement(c *fiber.Ctx) error
	ViewUsers(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	ViewOrders(c *fiber.Ctx) error
	DeliveryStatusUpdate(c *fiber.Ctx) error
	ManageUser(c *fiber.Ctx) error
	AdminRefresh(c *fiber.Ctx) error
}
