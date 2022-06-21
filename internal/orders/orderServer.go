package orders

type CreateOrderInput struct {
	Title           string
	CustomerOrderNo string
}

type OrderServer struct {
}

func (ordSev *OrderServer) CreateOrder(input *CreateOrderInput) {
	ordRepo := OrderRepo{}
	order := CreateOrderModel(input)
	ordRepo.CreateOrder(&order)
}
