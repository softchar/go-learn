package domain

type IOrderRepository interface {
	// CreateOrder 创建订单
	CreateOrder(order *Order)
}
