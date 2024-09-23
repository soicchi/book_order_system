package main

import (
	"log"

	"github.com/soicchi/book_order_system/config"
	"github.com/soicchi/book_order_system/presentation/router"

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
	router.NewRouter(r)
	
	r.Run()
}
