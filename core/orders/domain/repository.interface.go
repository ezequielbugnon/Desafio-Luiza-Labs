package domain

type IRepository interface {
	GetById(id string) (UserEntity, error)
	InsertFile(data []UserPresenter)
	GetByDate(startDate, endDate string) ([]UserEntity, error)
}

type IConnetionDb interface {
	Connect() error
	Disconnect()
	GetDB() interface{}
}
