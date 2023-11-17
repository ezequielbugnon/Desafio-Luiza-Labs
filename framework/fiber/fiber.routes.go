package framework

import (
	"github.com/ezequielbugnon/Desafio-Luiza-labs/core/orders/application"
	"github.com/gofiber/fiber/v2"
)

type fiberImplementation struct {
	fiber         *fiber.App
	ordersUseCase application.IOrdersUseCase
}

func New(f *fiber.App, o application.IOrdersUseCase) *fiberImplementation {
	return &fiberImplementation{
		fiber:         f,
		ordersUseCase: o,
	}
}

func (f *fiberImplementation) Routes() {
	api := f.fiber.Group("/api/v1")

	api.Post("/process-file", func(c *fiber.Ctx) error {
		return f.processFile(c)
	})

	api.Get("/user/:id", func(c *fiber.Ctx) error {
		return f.GetByID(c)
	})
}
