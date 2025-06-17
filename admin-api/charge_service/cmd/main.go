package main

import (
	handlers "charge_service/api"
	"charge_service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.CorsMiddleware)

	r.POST("/api/v1/charges", handlers.CreateTransaction)
	r.GET("/api/v1/charges", handlers.GetTransactions)

	r.Run(":7070")
}

