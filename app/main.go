package main

import (
	"log"

	"github.com/soicchi/book_order_system/config"
	"github.com/soicchi/book_order_system/presentation/router"

	"github.com/gin-gonic/gin"
	"github.com/soicchi/book_order_system/infrastructure/postgres"
)

func main() {
	// Load configuration
	config.LoadConfig()
	log.Println("Configuration loaded")

	// Database initialization
	postgres.Initialize()
	log.Println("Database initialized")

	r := gin.Default()
	router.NewRouter(r)

	r.Run()
}
