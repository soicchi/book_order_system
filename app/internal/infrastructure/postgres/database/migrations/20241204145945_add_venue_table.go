package migrations

import (
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var AddVenueTable = &gormigrate.Migration{
	ID: "20241204145945",
	Migrate: func(tx *gorm.DB) error {
		statement := `
			CREATE TABLE IF NOT EXISTS venues (
				id UUID PRIMARY KEY,
				name VARCHAR(255) NOT NULL,
				address VARCHAR(255) NOT NULL,
				capacity INT NOT NULL,
				created_at TIMESTAMP NOT NULL,
				updated_at TIMESTAMP NOT NULL
			);
		`

		if err := tx.Exec(statement).Error; err != nil {
			return fmt.Errorf("failed to create venues table: %w", err)
		}

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Migrator().DropTable("venues"); err != nil {
			return fmt.Errorf("failed to drop venues table: %w", err)
		}

		return nil
	},
}
