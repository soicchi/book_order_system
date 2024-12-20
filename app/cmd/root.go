package cmd

import (
	"os"
	"sync"

	"event_system/internal/config"
	"event_system/internal/infrastructure/postgres/database"
	"event_system/internal/logging"

	"github.com/spf13/cobra"
)

var (
	logger logging.Logger
	cfg    *config.Config
	once   sync.Once
)

var rootCmd = &cobra.Command{
	Use: "book_order_system",
}

func Execute() {
	once.Do(func() {
		// set up config
		cfg = config.LoadConfig()

		// Initialize logger
		logger = logging.InitLogger(cfg)

		// Connect to the database
		dbConfig := database.NewDBConfig(cfg)
		if err := dbConfig.Connect(); err != nil {
			panic(err)
		}
	})

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
