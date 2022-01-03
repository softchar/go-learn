package main

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/google/uuid"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Order struct {
	Id              string         `gorm:"column:Id;primaryKey;size:32"`
	CustomerOrderNo string         `gorm:"column:CustomerOrderNo;size:64"`
	OrderNo         sql.NullString `gorm:"column:OrderNo;size:64"`
	CustomerCode    sql.NullString `gorm:"column:CustomerCode;size:64"`
	CustomerName    sql.NullString `gorm:"column:CustomerName;size:64"`
	CreatedAt       time.Time      `gorm:"column:CreatedAt"`
	LadingTime      *time.Time     `gorm:"column:LadingTime"`
	Qutity          int            `gorm:"column:Qutity"`
	Volume          float32        `gorm:"column:Volume"`
}

func main() {

	db := connect_db()

	batch_create_orders(db)
}

// 链接数据库
func connect_db() *gorm.DB {
	dsn := "root:xiaozhang@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func create_order(db *gorm.DB) {
	order := Order{}
	order.Id = uuid.New().String()
	order.CustomerOrderNo = time_format_tostring(time.Now())
	db.Create(order)
}

func batch_create_orders(db *gorm.DB) {
	orders := make([]Order, 100)
	for i := 0; i < 100; i++ {
		order := Order{}
		order.Id = uuid.NewString()
		order.CustomerOrderNo = time_format_tostring(time.Now()) + strconv.Itoa(i)
		orders[i] = order
	}
	db.CreateInBatches(orders, 100)
}

func time_format_tostring(time time.Time) string {
	return time.Format("20060102150405") + strconv.Itoa(time.Nanosecond())
}
