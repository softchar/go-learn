package apps

import (
	"my/internal/routers"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouter(r *gin.Engine) *gin.Engine {
	r.POST("/order/create", routers.CreateOrder)
	return r
}
