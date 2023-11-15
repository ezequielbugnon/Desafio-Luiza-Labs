package application

type OrdersUseCase struct {
	repository string
}

func New(r string) *OrdersUseCase {
	return &OrdersUseCase{
		repository: r,
	}
}

func (o *OrdersUseCase) ProcessFile() string {
	return "hola"
}

func (o *OrdersUseCase) InsertFile() {

}
func (o *OrdersUseCase) RetrieveByID(id string) {

}
func (o *OrdersUseCase) RetrieveByPurchaseInterval(dateStart, dateEnd string) {

}
