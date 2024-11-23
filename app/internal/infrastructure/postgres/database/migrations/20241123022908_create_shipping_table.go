package migrations

import (
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var CreateShippingTable = &gormigrate.Migration{
	ID: "20241123022908",
	Migrate: func(tx *gorm.DB) error {
		statement := `
			CREATE TABLE IF NOT EXISTS shippings (
				id UUID PRIMARY KEY,
				order_id UUID NOT NULL,
				address VARCHAR(255) NOT NULL,
				method VARCHAR(255) NOT NULL,
				status VARCHAR(255) NOT NULL,
				created_at TIMESTAMP NOT NULL,
				updated_at TIMESTAMP NOT NULL
			);

			CREATE INDEX idx_shippings_order_id ON shippings (order_id);

			ALTER TABLE shippings ADD CONSTRAINT fk_shippings_order_id FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE;
		`

		if err := tx.Exec(statement).Error; err != nil {
			return fmt.Errorf("Error while creating shippings table: %v", err)
		}

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Migrator().DropTable("shippings"); err != nil {
			return fmt.Errorf("Error while dropping shippings table: %v", err)
		}

		return nil
	},
}
