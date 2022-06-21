package apps

import (
	"my/internal/orders"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {

	var createOrderInput orders.CreateOrderInput

	if err := c.ShouldBindJSON(&createOrderInput); err != nil {
		// 插入数据
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status:": "you are login"})
}
