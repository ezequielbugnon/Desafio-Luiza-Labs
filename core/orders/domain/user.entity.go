package domain

import "time"

type UserEntity struct {
	UserID int            `gorm:"column:usuario_id;primaryKey"`
	Name   string         `gorm:"column:nome"`
	Orders []OrdersEntity `gorm:"foreignKey:UserID"`
}

type OrdersEntity struct {
	OrderID  int                    `gorm:"column:pedido_id;primaryKey"`
	UserID   int                    `gorm:"column:usuario_id"`
	Products []ProductsOrdersEntity `gorm:"foreignKey:OrderID"`
}

type ProductsOrdersEntity struct {
	ID        int       `gorm:"column:produtos_pedido_id"`
	ProductID int       `gorm:"column:produto_id"`
	OrderID   int       `gorm:"column:pedido_id"`
	Value     float64   `gorm:"column:valor"`
	BuyDate   time.Time `gorm:"column:data_compra"`
}
