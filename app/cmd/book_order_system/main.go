package main

import (
	"github.com/soicchi/book_order_system/config"
	"github.com/soicchi/book_order_system/logger"
	"github.com/soicchi/book_order_system/infrastructure/postgres/database"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize logger
	logger := logger.InitLogger(cfg)
	logger.Info("Logger initialized")

	// Database initialization
	dbConfig := database.NewDBConnector(cfg)
	dbConfig.Connect()
	logger.Info("Database initialized")
}
