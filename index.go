package main

import (
	"my/internal/apps"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	apps.InitRouter(router)

	router.Run(":8080")
}
