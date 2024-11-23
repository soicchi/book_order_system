package migrations

import (
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var CreateProductTable = &gormigrate.Migration{
	ID: "20241123024240",
	Migrate: func(tx *gorm.DB) error {
		statement := `
			CREATE TABLE IF NOT EXISTS products (
				id UUID PRIMARY KEY,
				name VARCHAR(255) NOT NULL UNIQUE,
				price DECIMAL(10, 2) NOT NULL,
				created_at TIMESTAMP NOT NULL,
				updated_at TIMESTAMP NOT NULL
			);

			CREATE INDEX products_name_idx ON products (name);
		`

		if err := tx.Exec(statement).Error; err != nil {
			return fmt.Errorf("Error while creating products table: %v", err)
		}

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Migrator().DropTable("products"); err != nil {
			return fmt.Errorf("Error while dropping products table: %v", err)
		}

		return nil
	},
}
