package main

import (
	"github.com/soicchi/book_order_system/config"
	"github.com/soicchi/book_order_system/logger"
	"github.com/soicchi/book_order_system/presentation/router"

	"github.com/gin-gonic/gin"
	"github.com/soicchi/book_order_system/infrastructure/postgres"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize logger
	logger := logger.InitLogger(cfg)
	logger.Info("Logger initialized")

	// Database initialization
	postgres.Initialize(cfg)
	logger.Info("Database initialized")

	r := gin.Default()
	router.NewRouter(r, cfg)

	r.Run()
}
