package application

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ezequielbugnon/Desafio-Luiza-labs/core/orders/domain"
)

type OrdersUseCase struct {
	repository domain.IRepository
}

func New(r domain.IRepository) *OrdersUseCase {
	return &OrdersUseCase{
		repository: r,
	}
}

func (o *OrdersUseCase) ProcessFile(contentString string) ([]domain.UserPresenter, error) {

	var users []domain.UserPresenter

	scanner := bufio.NewScanner(strings.NewReader(contentString))
	for scanner.Scan() {
		line := scanner.Text()

		orders, err := parseLine(line)
		if err != nil {
			fmt.Println("Error al parsear la línea:", err)
			continue
		}

		var userExist *domain.UserPresenter
		for i, u := range users {
			if u.ID == orders.UserID {
				userExist = &users[i]
				break
			}
		}

		if userExist == nil {
			formattedDate := orders.BuyDate.Format("2006-01-02")
			userExist = &domain.UserPresenter{
				ID:   orders.UserID,
				Name: orders.Name,
				Orders: []domain.OrdersPresenter{
					{
						OrderID: orders.OrderID,
						Products: []domain.ProductPresenter{
							{
								ProductID: orders.ProductID,
								Value:     orders.Value,
								BuyDate:   formattedDate,
							},
						},
					},
				},
			}
			users = append(users, *userExist)
		} else {
			var ordersExist *domain.OrdersPresenter
			for i, p := range userExist.Orders {
				if p.OrderID == orders.OrderID {
					ordersExist = &userExist.Orders[i]
					break
				}
			}
			if ordersExist == nil {
				formattedDate := orders.BuyDate.Format("2006-01-02")
				insert := domain.OrdersPresenter{
					OrderID: orders.OrderID,
					Products: []domain.ProductPresenter{
						{
							ProductID: orders.ProductID,
							Value:     orders.Value,
							BuyDate:   formattedDate,
						},
					},
				}
				userExist.Orders = append(userExist.Orders, insert)
			} else {
				formattedDate := orders.BuyDate.Format("2006-01-02")
				insert := domain.ProductPresenter{
					ProductID: orders.ProductID,
					Value:     orders.Value,
					BuyDate:   formattedDate,
				}
				ordersExist.Products = append(ordersExist.Products, insert)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al escanear el archivo:", err)
		return []domain.UserPresenter{}, err
	}

	go o.repository.InsertFile(users)

	return users, nil
}

func (o *OrdersUseCase) RetrieveByID(id string) (domain.UserPresenter, error) {
	var userPresenter domain.UserPresenter
	userEntity, err := o.repository.GetById(id)
	if err != nil {
		return domain.UserPresenter{}, err
	}

	userPresenter.ID = userEntity.UserID
	userPresenter.Name = userEntity.Name

	for _, orderEntity := range userEntity.Orders {
		orderPresenter := domain.OrdersPresenter{
			OrderID: orderEntity.OrderID,
		}

		for _, productEntity := range orderEntity.Products {
			productPresenter := domain.ProductPresenter{
				ProductID: productEntity.ProductID,
				Value:     productEntity.Value,
				BuyDate:   productEntity.BuyDate.Format("2006-01-02"),
			}

			orderPresenter.Products = append(orderPresenter.Products, productPresenter)
		}

		userPresenter.Orders = append(userPresenter.Orders, orderPresenter)
	}

	return userPresenter, nil
}

func (o *OrdersUseCase) RetrieveByPurchaseInterval(dateStart, dateEnd time.Time) ([]domain.UserPresenter, error) {
	var userPresenters []domain.UserPresenter
	userEntity, err := o.repository.GetByDate(dateStart, dateEnd)
	if err != nil {
		return []domain.UserPresenter{}, err
	}

	for _, user := range userEntity {
		if !hasNonEmptyProducts(user.Orders) {
			continue
		}

		userPresenter := domain.UserPresenter{
			ID:     user.UserID,
			Name:   user.Name,
			Orders: make([]domain.OrdersPresenter, 0),
		}

		for _, orderEntity := range user.Orders {
			if len(orderEntity.Products) == 0 {
				continue
			}
			orderPresenter := domain.OrdersPresenter{
				OrderID:  orderEntity.OrderID,
				Products: make([]domain.ProductPresenter, 0),
			}

			for _, productEntity := range orderEntity.Products {
				productPresenter := domain.ProductPresenter{
					ProductID: productEntity.ProductID,
					Value:     productEntity.Value,
					BuyDate:   productEntity.BuyDate.Format("2006-01-02"),
				}

				orderPresenter.Products = append(orderPresenter.Products, productPresenter)
			}

			userPresenter.Orders = append(userPresenter.Orders, orderPresenter)
		}
		userPresenters = append(userPresenters, userPresenter)
	}
	return userPresenters, nil

}

func parseLine(line string) (OrderParser, error) {
	fields := []string{
		strings.TrimSpace(line[0:10]),
		strings.TrimSpace(line[10:55]),
		strings.TrimSpace(line[55:65]),
		strings.TrimSpace(line[65:75]),
		strings.TrimSpace(line[75:87]),
		strings.TrimSpace(line[87:]),
	}

	userID, err := strconv.Atoi(fields[0])
	if err != nil {
		return OrderParser{}, err
	}

	name := fields[1]

	orderID, err := strconv.Atoi(fields[2])
	if err != nil {
		return OrderParser{}, err
	}

	productID, err := strconv.Atoi(fields[3])
	if err != nil {
		return OrderParser{}, err
	}

	value, err := strconv.ParseFloat(fields[4], 64)
	if err != nil {
		return OrderParser{}, err
	}

	buyDate, err := time.Parse("20060102", fields[5])
	if err != nil {
		return OrderParser{}, err
	}

	return OrderParser{
		OrderID:   orderID,
		Name:      name,
		UserID:    userID,
		ProductID: productID,
		Value:     value,
		BuyDate:   buyDate,
	}, nil
}

func hasNonEmptyProducts(orders []domain.OrdersEntity) bool {
	for _, order := range orders {
		if len(order.Products) > 0 {
			return true
		}
	}
	return false
}
