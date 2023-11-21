package application

import (
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/ezequielbugnon/Desafio-Luiza-labs/core/orders/domain"
	"github.com/stretchr/testify/assert"
)

type MockRepository struct {
	Users map[string]domain.UserEntity
}

func (m *MockRepository) InsertFile(users []domain.UserPresenter) {
	for _, user := range users {
		userEntity := domain.UserEntity{
			UserID: user.ID,
			Name:   user.Name,
			Orders: make([]domain.OrdersEntity, len(user.Orders)),
		}

		for i, order := range user.Orders {
			orderEntity := domain.OrdersEntity{
				OrderID:  order.OrderID,
				Products: make([]domain.ProductsOrdersEntity, len(order.Products)),
			}

			for j, product := range order.Products {
				productEntity := domain.ProductsOrdersEntity{
					ProductID: product.ProductID,
					Value:     product.Value,
					BuyDate:   time.Time{},
				}

				productEntity.BuyDate, _ = time.Parse("2006-01-02", product.BuyDate)

				orderEntity.Products[j] = productEntity
			}

			userEntity.Orders[i] = orderEntity
		}

		m.Users[strconv.Itoa(userEntity.UserID)] = userEntity
	}
}

func (m *MockRepository) GetById(id string) (domain.UserEntity, error) {
	user, ok := m.Users[id]
	if !ok {
		return domain.UserEntity{}, errors.New("err GetById repository")
	}
	return user, nil
}

func (m *MockRepository) GetByDate(dateStart, dateEnd time.Time) ([]domain.UserEntity, error) {
	var users []domain.UserEntity
	for _, user := range m.Users {
		users = append(users, user)
	}
	return users, nil
}

func TestOrdersUseCase_ProcessFile(t *testing.T) {
	contentString := "0000000070                              Palmer Prosacco00000007530000000003     1836.7420210308"

	mockRepository := &MockRepository{
		Users: make(map[string]domain.UserEntity),
	}

	useCase := New(mockRepository)

	users, err := useCase.ProcessFile(contentString)

	assert.NoError(t, err)

	assert.NotEmpty(t, users)
}

func TestOrdersUseCase_ProcessFile_WithSameID(t *testing.T) {
	contentString := "0000000070                              Palmer Prosacco00000007530000000003     1836.7420210308\n" +
		"0000000070                              Palmer Prosacco00000007530000000004      618.7920210308\n"

	mockRepository := &MockRepository{
		Users: make(map[string]domain.UserEntity),
	}

	useCase := New(mockRepository)

	users, err := useCase.ProcessFile(contentString)

	assert.NoError(t, err)

	assert.NotEmpty(t, users)
}

func TestOrdersUseCase_ProcessFileWithErr(t *testing.T) {
	contentString := "0000000070                              Palmer Prosacco00000007530000000003    1836.7420210308ggxxxx"

	mockRepository := &MockRepository{
		Users: make(map[string]domain.UserEntity),
	}

	useCase := New(mockRepository)

	_, err := useCase.ProcessFile(contentString)
	println(err)

	assert.Nil(t, err)
}

func TestOrdersUseCase_RetrieveByID(t *testing.T) {
	mockRepository := &MockRepository{
		Users: map[string]domain.UserEntity{
			"1": {
				UserID: 1,
				Name:   "John Doe",
				Orders: []domain.OrdersEntity{
					{
						OrderID:  101,
						Products: []domain.ProductsOrdersEntity{},
					},
				},
			},
		},
	}

	useCase := New(mockRepository)

	user, err := useCase.RetrieveByID("1")

	assert.NoError(t, err)

	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "John Doe", user.Name)
}

func TestOrdersUseCase_RetrieveByPurchaseInterval(t *testing.T) {

	mockRepository := &MockRepository{
		Users: map[string]domain.UserEntity{
			"1": {
				UserID: 1,
				Name:   "John Doe",
				Orders: []domain.OrdersEntity{
					{
						OrderID: 101,
						Products: []domain.ProductsOrdersEntity{
							{
								ProductID: 201,
								Value:     50.0,
								BuyDate:   time.Now(),
							},
						},
					},
				},
			},
		},
	}

	useCase := New(mockRepository)

	dateStart := time.Now().AddDate(0, -1, 0)
	dateEnd := time.Now()

	users, err := useCase.RetrieveByPurchaseInterval(dateStart, dateEnd)

	assert.NoError(t, err)

	assert.NotEmpty(t, users)
}
