package application

import (
	"time"

	"github.com/ezequielbugnon/Desafio-Luiza-labs/core/orders/domain"
)

type IOrdersUseCase interface {
	ProcessFile(contentString string) ([]domain.UserPresenter, error)
	RetrieveByID(id string) (domain.UserPresenter, error)
	RetrieveByPurchaseInterval(dateStart, dateEnd time.Time) ([]domain.UserPresenter, error)
}
