package framework

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (f *fiberImplementation) processFile(c *fiber.Ctx) error {
	file, err := c.FormFile("data")
	if err != nil {
		fmt.Println("Erro ao obter arquivo:", err)
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Envie um arquivo .txt com o formato correspondente e o valor “data” "})
	}

	fileContent, err := file.Open()
	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer fileContent.Close()

	contentBytes, err := io.ReadAll(fileContent)
	if err != nil {
		fmt.Println("Erro ao ler o conteúdo do arquivo:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	contentString := string(contentBytes)

	result, err := f.ordersUseCase.ProcessFile(contentString)
	if err != nil {
		fmt.Println("Erro ao ler o conteúdo do arquivo:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func (f *fiberImplementation) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := f.ordersUseCase.RetrieveByID(id)
	if err != nil {
		log.Println("ocorreu um erro", err)
		if strings.Contains(err.Error(), "Nenhum usuário encontrado com ID") {
			return c.Status(fiber.StatusNotFound).JSON(err.Error())
		}
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
