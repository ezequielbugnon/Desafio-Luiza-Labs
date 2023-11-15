package framework

import (
	"github.com/ezequielbugnon/Desafio-Luiza-labs/core/orders/application"
	"github.com/gofiber/fiber/v2"
)

type fiberImplemantation struct {
	fiber         *fiber.App
	ordersUseCase application.IOrdersUseCase
}

func New(f *fiber.App, o application.IOrdersUseCase) *fiberImplemantation {
	return &fiberImplemantation{
		fiber:         f,
		ordersUseCase: o,
	}
}

func (f *fiberImplemantation) Routes() {
	api := f.fiber.Group("/api/v1")

	api.Post("/process-file", func(c *fiber.Ctx) error {
		return f.processFile(c)
	})
}
