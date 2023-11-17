package domain

import "time"

type IRepository interface {
	GetById(id string) (UserEntity, error)
	InsertFile(data []UserPresenter)
	GetByDate(startDate, endDate time.Time) ([]UserEntity, error)
}

type IConnetionDb interface {
	Connect() error
	Disconnect()
	GetDB() interface{}
}
