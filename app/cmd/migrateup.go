package cmd

import (
	"log/slog"

	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/database"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrateup",
	Short: "Migrate the database to the latest version",
	Long: `This command migrates the database to the latest revision available. 
It runs all the necessary migrations sequentially, ensuring that the database 
schema is up-to-date and matches the latest application requirements. By executing 
this command, you can apply new migrations, update existing tables, add or modify 
indexes, constraints, and other database objects as specified in the migration files.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Migrating the database...")

		if err := database.Migrate(); err != nil {
			logger.Error("Failed to migrate the database", slog.Any("error", err))
			return
		}

		logger.Info("Database migrated successfully")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
