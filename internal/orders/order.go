package orders

import "time"

type Order struct {
	Id              int       `gorm:"column:Id;primaryKey;size:11"`
	Title           string    `gorm:"column:Title;size:64"`
	CustomerOrderNo string    `gorm:"column:CustomerOrderNo;size:125"`
	OrderNo         string    `gorm:"column:OrderNo;size:64"`
	IsDeleted       bool      `gorm:"column:IsDeleted;size:1"`
	CreateTime      time.Time `gorm:"column:CreateTime;size:6"`
}

func CreateOrderModel(input *CreateOrderInput) Order {
	return Order{
		Title:           input.Title,
		CustomerOrderNo: input.CustomerOrderNo,
		OrderNo:         "",
		IsDeleted:       false,
		CreateTime:      time.Now(),
	}
}
