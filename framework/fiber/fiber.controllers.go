package framework

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Usuario struct {
	ID      int               `json:"id"`
	Nome    string            `json:"nome"`
	Pedidos []PedidoPresenter `json:"pedidos"`
}

type Pedido struct {
	UsuarioID  int       `json:"usuario_id"`
	Nome       string    `json:"nome"`
	PedidoID   int       `json:"pedido_id"`
	ProdutoID  int       `json:"produto_id"`
	Valor      float64   `json:"valor"`
	DataCompra time.Time `json:"data_compra"`
}

type PedidoPresenter struct {
	PedidoID   int     `json:"pedido_id"`
	ProdutoID  int     `json:"produto_id"`
	Valor      float64 `json:"valor"`
	DataCompra string  `json:"data_compra"`
}

func (f *fiberImplemantation) processFile(c *fiber.Ctx) error {
	file, err := c.FormFile("text")
	if err != nil {
		fmt.Println("Error al obtener el archivo:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	fileContent, err := file.Open()
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer fileContent.Close()

	contentBytes, err := io.ReadAll(fileContent)
	if err != nil {
		fmt.Println("Error al leer el contenido del archivo:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	contentString := string(contentBytes)

	fmt.Println(f.ordersUseCase.ProcessFile())

	var usuarios []Usuario

	scanner := bufio.NewScanner(strings.NewReader(contentString))
	for scanner.Scan() {
		line := scanner.Text()

		pedido, err := parseLine(line)
		if err != nil {
			fmt.Println("Error al parsear la línea:", err)
			continue
		}

		var usuarioExistente *Usuario
		for i, u := range usuarios {
			if u.ID == pedido.UsuarioID {
				usuarioExistente = &usuarios[i]
				break
			}
		}

		if usuarioExistente == nil {
			formattedDate := pedido.DataCompra.Format("2006-01-02")
			usuarioExistente = &Usuario{
				ID:   pedido.UsuarioID,
				Nome: pedido.Nome,
				Pedidos: []PedidoPresenter{
					{
						PedidoID:   pedido.PedidoID,
						ProdutoID:  pedido.ProdutoID,
						Valor:      pedido.Valor,
						DataCompra: formattedDate,
					},
				},
			}
			usuarios = append(usuarios, *usuarioExistente)
		} else {
			formattedDate := pedido.DataCompra.Format("2006-01-02")
			insert := PedidoPresenter{
				PedidoID:   pedido.PedidoID,
				ProdutoID:  pedido.ProdutoID,
				Valor:      pedido.Valor,
				DataCompra: formattedDate,
			}
			usuarioExistente.Pedidos = append(usuarioExistente.Pedidos, insert)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al escanear el archivo:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(usuarios)
}

func parseLine(line string) (Pedido, error) {
	fields := []string{
		strings.TrimSpace(line[0:10]),
		strings.TrimSpace(line[10:55]),
		strings.TrimSpace(line[55:65]),
		strings.TrimSpace(line[65:75]),
		strings.TrimSpace(line[75:87]),
		strings.TrimSpace(line[87:]),
	}

	usuarioID, err := strconv.Atoi(fields[0])
	if err != nil {
		return Pedido{}, err
	}

	nome := fields[1]

	pedidoID, err := strconv.Atoi(fields[2])
	if err != nil {
		return Pedido{}, err
	}

	productoID, err := strconv.Atoi(fields[3])
	if err != nil {
		return Pedido{}, err
	}

	valor, err := strconv.ParseFloat(fields[4], 64)
	if err != nil {
		return Pedido{}, err
	}

	dataCompra, err := time.Parse("20060102", fields[5])
	if err != nil {
		return Pedido{}, err
	}

	return Pedido{
		PedidoID:   pedidoID,
		Nome:       nome,
		UsuarioID:  usuarioID,
		ProdutoID:  productoID,
		Valor:      valor,
		DataCompra: dataCompra,
	}, nil
}
