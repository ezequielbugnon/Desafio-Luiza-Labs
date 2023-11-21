package framework

import (
	"fmt"
	"io"
	"strings"
	"time"

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
		if strings.Contains(err.Error(), "Nenhum usuário encontrado com ID") {
			return c.Status(fiber.StatusNotFound).JSON(err.Error())
		}
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func (f *fiberImplementation) GetByDate(c *fiber.Ctx) error {
	start := c.Query("start")
	end := c.Query("end")

	startTime, err := parseDate(start)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Formato de data inválido"})
	}

	endTime, err := parseDate(end)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Formato de data inválido"})
	}

	result, err := f.ordersUseCase.RetrieveByPurchaseInterval(startTime, endTime)
	if err != nil {

		if strings.Contains(err.Error(), "Nenhum usuário encontrado") {
			return c.Status(fiber.StatusNotFound).JSON(err.Error())
		}
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func parseDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}
