package main

import (
	"fmt"

	"github.com/ezequielbugnon/Desafio-Luiza-labs/core/orders/application"
	"github.com/ezequielbugnon/Desafio-Luiza-labs/core/orders/domain"
	"github.com/ezequielbugnon/Desafio-Luiza-labs/core/orders/infraestructure"
	"github.com/ezequielbugnon/Desafio-Luiza-labs/database"
	framework "github.com/ezequielbugnon/Desafio-Luiza-labs/framework/fiber"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	dns := "user=postgres dbname=postgres password=postgres host=localhost port=5432 sslmode=disable"
	connection, err := database.NewPostgres(dns)
	if err != nil {
		fmt.Printf("error al conectar la base de datos : %v", err)
	}

	err = connection.Database.AutoMigrate(&domain.UserEntity{}, &domain.OrdersEntity{}, &domain.ProductsOrdersEntity{})
	if err != nil {
		fmt.Printf("error al migrar datos : %v", err)
	}
	repository := infraestructure.New(connection)
	ordersUseCase := application.New(repository)

	routes := framework.New(app, ordersUseCase)

	routes.Routes()

	err = app.Listen(":3000")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
