package main

import (
	"log-service/config"
	"log-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// DB 연결
	config.ConnectDB()

	// 라우터 등록
	routes.RegisterRoutes(r)

	r.Run(":8080")
}
