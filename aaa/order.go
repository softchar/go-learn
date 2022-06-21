package aaa

func CreateOrder() {

}

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"strconv"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// func main() {

// }

// // 链接数据库
// func connect_db() *gorm.DB {
// 	dsn := "root:xiaozhang@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	return db
// }

// func querys(db *gorm.DB) {
// 	orders := []models.Order{}

// 	// db.Find(&orders)
// 	db.Where(models.Order{OrderType: 50}).Find(&orders)

// 	fmt.Println(len(orders))

// 	for _, v := range orders {
// 		fmt.Println(v)
// 	}
// }

// func create_order(db *gorm.DB) {
// 	order := models.Order{}
// 	order.Id = uuid.New().String()
// 	order.CustomerOrderNo = time_format_tostring(time.Now())
// 	db.Create(order)
// }

// func batch_create_orders(db *gorm.DB) {
// 	orders := make([]models.Order, 100)
// 	for i := 0; i < 100; i++ {
// 		order := models.Order{}
// 		order.Id = uuid.NewString()
// 		order.CustomerOrderNo = time_format_tostring(time.Now()) + strconv.Itoa(i)
// 		order.CreatedAt = time.Now()
// 		order.OrderType = rand.Intn(100)
// 		orders[i] = order
// 	}
// 	db.CreateInBatches(orders, 100)
// }

// func time_format_tostring(time time.Time) string {
// 	return time.Format("20060102150405") + strconv.Itoa(time.Nanosecond())
// }
