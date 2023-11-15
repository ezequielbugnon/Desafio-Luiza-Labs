package main

import (
	"fmt"

	framework "github.com/ezequielbugnon/Desafio-Luiza-labs/framework/fiber"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes := framework.New(app)

	routes.Routes()

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}

}
