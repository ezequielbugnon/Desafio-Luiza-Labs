package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ezequielbugnon/Desafio-Luiza-labs/core/orders/application"
	"github.com/ezequielbugnon/Desafio-Luiza-labs/core/orders/domain"
	"github.com/ezequielbugnon/Desafio-Luiza-labs/core/orders/infraestructure"
	"github.com/ezequielbugnon/Desafio-Luiza-labs/database"
	framework "github.com/ezequielbugnon/Desafio-Luiza-labs/framework/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost, http://127.0.0.1",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}
	dns := os.Getenv("DNS")
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{TranslateError: true})
	if err != nil {
		fmt.Printf("erro ao conectar o banco de dados: %v", err)
	}

	connection, err := database.NewPostgres(db)
	if err != nil {
		fmt.Printf("erro ao conectar o banco de dados: %v", err)
	}

	err = connection.Database.AutoMigrate(&domain.UserEntity{}, &domain.OrdersEntity{}, &domain.ProductsOrdersEntity{})
	if err != nil {
		fmt.Printf("error ao migrar dados : %v", err)
	}
	repository := infraestructure.New(connection)
	ordersUseCase := application.New(repository)

	routes := framework.New(app, ordersUseCase)

	routes.Routes()

	err = app.Listen(":8080")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
