package migrations

import (
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var CreateUserTable = &gormigrate.Migration{
	ID: "20241124101755",
	Migrate: func(tx *gorm.DB) error {
		statement := `
			CREATE TABLE IF NOT EXISTS users (
				id UUID PRIMARY KEY,
				username VARCHAR(255) NOT NULL,
				email VARCHAR(255) NOT NULL UNIQUE,
				password VARCHAR(255) NOT NULL,
				created_at TIMESTAMP NOT NULL,
				updated_at TIMESTAMP NOT NULL
			);
		`

		if err := tx.Exec(statement).Error; err != nil {
			return fmt.Errorf("failed to create users table: %w", err)
		}

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Migrator().DropTable("users"); err != nil {
			return fmt.Errorf("failed to drop users table: %w", err)
		}

		return nil
	},
}
