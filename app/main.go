package main

import (
	"log"

	"github.com/soicchi/book_order_system/infrastructure/postgres"

	"github.com/soicchi/book_order_system/config"
)

func main() {
	// Load configuration
	config.LoadConfig()
	log.Println("Configuration loaded")


	// Database initialization
	postgres.Initialize()
	log.Println("Database initialized")
}
