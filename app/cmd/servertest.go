package cmd

import (
	"event_system/internal/infrastructure/postgres/database"
	"event_system/internal/presentation/router"
	"event_system/internal/presentation/validator"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

// servertestCmd represents the servertest command
var servertestCmd = &cobra.Command{
	Use:   "servertest",
	Short: "Start the test API server",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		dbConfig := database.NewTestDBConfig()

		if err := dbConfig.Connect(); err != nil {
			logger.Error("Failed to connect to the database", err)
			panic(err)
		}

		logger.Info("Database initialized")

		e := echo.New()
		e.Validator = validator.NewCustomValidator()

		router.NewRouter(e, logger)

		e.Logger.Fatal(e.Start(":8081"))
	},
}

func init() {
	rootCmd.AddCommand(servertestCmd)
}
