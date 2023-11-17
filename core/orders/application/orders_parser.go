package application

import "time"

type OrderParser struct {
	UserID    int       `json:"usuario_id"`
	Name      string    `json:"nome"`
	OrderID   int       `json:"pedido_id"`
	ProductID int       `json:"produto_id"`
	Value     float64   `json:"valor"`
	BuyDate   time.Time `json:"data_compra"`
}
