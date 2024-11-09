package main

import (
	"github.com/soicchi/book_order_system/config"
	"github.com/soicchi/book_order_system/infrastructure/postgres/database"
	"github.com/soicchi/book_order_system/logger"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize logger
	logger := logger.InitLogger(cfg)
	logger.Info("Logger initialized")

	// Database initialization
	dbConfig := database.NewDBConfig(cfg)
	dbConfig.Connect()
	logger.Info("Database initialized")
}
