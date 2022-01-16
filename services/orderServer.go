package services

import (
	"my.io/services/Inputs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderServer struct {

}

// CreateOrder 创建订单
func (orderServer *OrderServer) CreateOrder(c *gin.Context) {

	var createOrderInput Inputs.CreateOrderInput

	if err := c.ShouldBindJSON(&createOrderInput); err != nil {
		// 插入数据
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status:": "you are login"})
}