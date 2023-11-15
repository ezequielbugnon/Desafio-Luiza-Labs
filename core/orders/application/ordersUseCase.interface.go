package application

type IOrdersUseCase interface {
	ProcessFile() string
	InsertFile()
	RetrieveByID(id string)
	RetrieveByPurchaseInterval(dateStart, dateEnd string)
}
