package orders

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type OrderRepo struct{}

func (ordRepo *OrderRepo) CreateOrder(order *Order) 

	dsn := "root:xiaozhang@tcp(127.0.0.1:3306)/gotest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.Create(order)
}

func GetOrders() {

}
