package interfaces

import "github.com/gofiber/fiber/v2"

type IProduct interface {
	AddProducts(*fiber.Ctx) error
	UpdatePro(*fiber.Ctx) error
	DelProduct(*fiber.Ctx) error
	ViewProducts(*fiber.Ctx) error
	GetbyCategory(*fiber.Ctx) error
	SearchProduct(*fiber.Ctx) error
}
