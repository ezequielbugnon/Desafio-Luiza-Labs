package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Usuario struct {
	ID      int      `json:"id"`
	Nome    string   `json:"nome"`
	Pedidos []Pedido `json:"pedidos"`
}

type Pedido struct {
	UsuarioID  int       `json:"usuario_id"`
	Nome       string    `json:"nome"`
	PedidoID   int       `json:"pedido_id"`
	ProdutoID  int       `json:"produto_id"`
	Valor      float64   `json:"valor"`
	DataCompra time.Time `json:"data_compra"`
}

func main() {
	filePath := "cmd/data_1.txt" // Reemplaza con la ruta de tu archivo

	// Abre el archivo
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	var usuarios []Usuario

	// Lee el archivo línea por línea
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println("Línea:", line)

		// Parsea la línea y crea un Pedido
		pedido, err := parseLine(line)
		if err != nil {
			fmt.Println("Error al parsear la línea:", err)
			continue
		}

		// Busca si el usuario ya existe en la lista de usuarios
		var usuarioExistente *Usuario
		for i, u := range usuarios {
			if u.ID == pedido.UsuarioID {
				usuarioExistente = &usuarios[i]
				break
			}
		}

		// Si el usuario no existe, crea uno nuevo y agrégalo a la lista
		if usuarioExistente == nil {
			usuarioExistente = &Usuario{
				ID:      pedido.UsuarioID,
				Nome:    pedido.Nome,
				Pedidos: []Pedido{pedido},
			}
			usuarios = append(usuarios, *usuarioExistente)
		} else {
			// Si el usuario ya existe, agrega el pedido a su lista de pedidos
			usuarioExistente.Pedidos = append(usuarioExistente.Pedidos, pedido)
		}
	}

	// Convierte los usuarios a formato JSON
	jsonData, err := json.Marshal(usuarios)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	// Imprime el JSON resultante
	fmt.Println(string(jsonData))
}

// La función parseLine
func parseLine(line string) (Pedido, error) {
	// Asumiendo que cada línea tiene un formato específico
	fields := []string{
		strings.TrimSpace(line[0:10]),  // ID del Pedido
		strings.TrimSpace(line[10:55]), // Nombre del Usuario
		strings.TrimSpace(line[55:65]),
		strings.TrimSpace(line[65:75]),
		strings.TrimSpace(line[75:87]), // Valor
		strings.TrimSpace(line[87:]),   // Fecha de Compra
	}

	// Convierte los campos a los tipos correspondientes
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

	// Crea y retorna un Pedido
	return Pedido{
		PedidoID:   pedidoID,
		Nome:       nome,
		UsuarioID:  usuarioID,
		ProdutoID:  productoID,
		Valor:      valor,
		DataCompra: dataCompra,
	}, nil
}
