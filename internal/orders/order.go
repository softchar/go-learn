package orders

import "time"

type Order struct {
	Id              string     `gorm:"column:Id;primaryKey;size:36"`
	CustomerOrderNo string     `gorm:"column:CustomerOrderNo;size:64"`
	OrderNo         string     `gorm:"column:OrderNo;size:64"`
	CustomerCode    string     `gorm:"column:CustomerCode;size:64"`
	CustomerName    string     `gorm:"column:CustomerName;size:64"`
	CreatedAt       time.Time  `gorm:"column:CreatedAt"`
	LadingTime      *time.Time `gorm:"column:LadingTime"`
	Qutity          int        `gorm:"column:Qutity"`
	Volume          float32    `gorm:"column:Volume"`
	OrderType       int        `gorm:"column:OrderType"`
}
