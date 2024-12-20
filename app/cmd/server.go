package cmd

import (
	"event_system/internal/infrastructure/postgres/database"
	"event_system/internal/presentation/router"
	"event_system/internal/presentation/validator"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

// APIサーバーを起動するためのコマンドを定義
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
		// Database initialization
		dbConfig := database.NewDBConfig(cfg)

		// Connect to the database
		if err := dbConfig.Connect(); err != nil {
			logger.Error("Failed to connect to the database", err)
			panic(err)
		}

		logger.Info("Database initialized")

		// Initialize Echo
		e := echo.New()

		// resister validator
		e.Validator = validator.NewCustomValidator()

		// set up routers
		router.NewRouter(e, logger)

		// Output all routes in local
		if cfg.Environment == "local" {
			router.OutputRoutes(e)
		}

		e.Logger.Fatal(e.Start(":8080"))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
