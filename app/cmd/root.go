package cmd

import (
	"os"
	"sync"

	"github.com/soicchi/book_order_system/config"
	"github.com/soicchi/book_order_system/infrastructure/postgres/database"
	loggerPkg "github.com/soicchi/book_order_system/logger" // To avoid conflict with logger variable name

	"github.com/spf13/cobra"
)

var (
	logger loggerPkg.Logger
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

		// set up logger
		logger = loggerPkg.InitLogger(cfg)

		// Connect to the database
		dbConfig := database.NewDBConfig(cfg)
		dbConfig.Connect()
	})

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
