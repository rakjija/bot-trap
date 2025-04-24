package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rakjija/bot-trap/observer/internal/routes"
)

func main() {
	r := gin.Default()

	r.POST("/logs", routes.HandleLog)

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.Run(":9000")

	fmt.Println("[INFO] Server is running on 9000...")
}
