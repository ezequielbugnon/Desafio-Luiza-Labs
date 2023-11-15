package framework

import "github.com/gofiber/fiber/v2"

type fiberImplemantation struct {
	fiber *fiber.App
}

func New(f *fiber.App) *fiberImplemantation {
	return &fiberImplemantation{
		fiber: f,
	}
}

func (f *fiberImplemantation) Routes() {
	api := f.fiber.Group("/api/v1")

	api.Post("/process-file", processFile)
}
