package domain

type UserPresenter struct {
	ID     int               `json:"id"`
	Name   string            `json:"nome"`
	Orders []OrdersPresenter `json:"pedidos"`
}

type OrdersPresenter struct {
	OrderID  int                `json:"pedido_id"`
	Products []ProductPresenter `json:"produtos"`
}

type ProductPresenter struct {
	ProductID int     `json:"produto_id"`
	Value     float64 `json:"valor"`
	BuyDate   string  `json:"data_compra"`
}
