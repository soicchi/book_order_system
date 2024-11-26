package migrations

import (
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var AddColumnStatusInOrderTable = &gormigrate.Migration{
	ID: "20241125234339",
	Migrate: func(tx *gorm.DB) error {
		statement := `
			ALTER TABLE orders ADD COLUMN status VARCHAR(255) NOT NULL;
		`

		if err := tx.Exec(statement).Error; err != nil {
			return fmt.Errorf("failed to add column status in orders table: %w", err)
		}

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Migrator().DropColumn("orders", "status"); err != nil {
			return fmt.Errorf("failed to drop column status in orders table: %w", err)
		}

		return nil
	},
}
