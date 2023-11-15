package main

import (
	"fmt"

	"github.com/ezequielbugnon/Desafio-Luiza-labs/core/orders/application"
	framework "github.com/ezequielbugnon/Desafio-Luiza-labs/framework/fiber"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	ordersUseCase := application.New("repo")

	routes := framework.New(app, ordersUseCase)

	routes.Routes()

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
