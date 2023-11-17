package infraestructure

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ezequielbugnon/Desafio-Luiza-labs/core/orders/domain"
	"gorm.io/gorm"
)

type GormRepository struct {
	Db domain.IConnetionDb
}

func New(database domain.IConnetionDb) *GormRepository {
	return &GormRepository{
		Db: database,
	}
}
func (g *GormRepository) GetById(id string) (domain.UserEntity, error) {
	db := g.Db.GetDB().(*gorm.DB)
	var user domain.UserEntity

	result := db.Preload("Orders.Products").Where("usuario_id = ?", id).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, fmt.Errorf("Nenhum usuário encontrado com ID %s", id)
		}
		log.Println("Erro em gorm -> getbyid", result.Error)

		return user, fmt.Errorf("Ocorreu um erro de ID de comunicação %s", id)
	}

	return user, nil
}

func (g *GormRepository) InsertFile(data []domain.UserPresenter) {
	db := g.Db.GetDB().(*gorm.DB)

	var userEntities []domain.UserEntity

	for _, user := range data {
		userEntity := domain.UserEntity{
			UserID: user.ID,
			Name:   user.Name,
			Orders: make([]domain.OrdersEntity, len(user.Orders)),
		}

		for i, order := range user.Orders {
			orderEntity := domain.OrdersEntity{
				OrderID:  order.OrderID,
				UserID:   user.ID,
				Products: make([]domain.ProductsOrdersEntity, len(order.Products)),
			}

			for j, product := range order.Products {
				buyDate, _ := time.Parse("2006-01-02", product.BuyDate)

				productEntity := domain.ProductsOrdersEntity{
					ProductID: product.ProductID,
					OrderID:   order.OrderID,
					Value:     product.Value,
					BuyDate:   buyDate,
				}

				orderEntity.Products[j] = productEntity
			}

			userEntity.Orders[i] = orderEntity
		}

		userEntities = append(userEntities, userEntity)
	}

	result := db.Save(userEntities)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("Ok")
}

func (g *GormRepository) GetByDate(startDate, endDate string) ([]domain.UserEntity, error) {
	db := g.Db.GetDB().(*gorm.DB)

	var userEntities []domain.UserEntity

	result := db.Preload("Orders.Products").Joins("JOIN orders_entities ON orders_entities.usuario_id = usuario_entities.usuario_id").
		Joins("JOIN produtos_pedidos_entities ON produtos_pedidos_entities.pedido_id = orders_entities.pedido_id").
		Where("produtos_pedidos_entities.data_compra BETWEEN ? AND ?", startDate, endDate).
		Find(&userEntities)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		log.Println("Nenhum resultado foi encontrado para os critérios especificados.")
	}

	return userEntities, nil
}
