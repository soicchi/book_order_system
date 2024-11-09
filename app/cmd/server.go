package cmd

import (
	"github.com/soicchi/book_order_system/config"
	"github.com/soicchi/book_order_system/infrastructure/postgres/database"
	"github.com/soicchi/book_order_system/logger"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the API server",
	Long: `Launches the API server built with Echo, a high-performance, extensible,
and minimalist web framework in Go. The server provides a RESTful API for managing
customer data, orders, books, and other resources. 

This command starts the server on the specified port, allowing clients to send HTTP
requests to interact with the application's core functionalities.

Use this command to initialize the server and make the application accessible to 
clients, or to test the API endpoints directly.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load configuration
		cfg := config.LoadConfig()

		// Initialize logger
		logger := logger.InitLogger(cfg)
		logger.Info("Logger initialized")

		// Database initialization
		dbConfig := database.NewDBConfig(cfg)
		dbConfig.Connect()
		logger.Info("Database initialized")

		// Initialize Echo
		e := echo.New()
		e.GET("/", func(c echo.Context) error {
			return c.String(200, "Hello, World!")
		})

		e.Logger.Fatal(e.Start(":8080"))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
