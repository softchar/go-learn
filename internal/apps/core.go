package apps

import (
	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouter(r *gin.Engine) *gin.Engine {
	r.POST("/order/create", CreateOrder)
	return r
}
