package migrations

import (
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var CreateCustomerTable = &gormigrate.Migration{
	ID: "20241122054931",
	Migrate: func(tx *gorm.DB) error {
		statement := `
			CREATE TABLE IF NOT EXISTS customers (
				id UUID PRIMARY KEY,
				name VARCHAR(255) NOT NULL,
				email VARCHAR(255) NOT NULL UNIQUE,
				created_at TIMESTAMP NOT NULL,
				updated_at TIMESTAMP NOT NULL
			)
		`

		if err := tx.Exec(statement).Error; err != nil {
			return fmt.Errorf("Error while creating customers table: %v", err)
		}

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Migrator().DropTable("customers"); err != nil {
			return fmt.Errorf("Error while dropping customers table: %v", err)
		}

		return nil
	},
}
