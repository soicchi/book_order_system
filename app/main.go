package main

import (
	"log"
	"net/http"

	"github.com/soicchi/book_order_system/config"

	"github.com/soicchi/book_order_system/infrastructure/postgres"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.LoadConfig()
	log.Println("Configuration loaded")


	// Database initialization
	postgres.Initialize()
	log.Println("Database initialized")

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})
	r.Run()
}
