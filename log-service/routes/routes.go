package routes

import (
	"log-service/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", handlers.Ping)
	r.POST("/log", handlers.SaveLog)
	r.GET("/logs", handlers.GetLogs)
}
