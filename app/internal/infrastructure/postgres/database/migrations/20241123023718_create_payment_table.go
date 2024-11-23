package migrations

import (
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var CreatePaymentTable = &gormigrate.Migration{
	ID: "20241123023718",
	Migrate: func(tx *gorm.DB) error {
		statement := `
			CREATE TABLE IF NOT EXISTS payments (
				id UUID PRIMARY KEY,
				order_id UUID NOT NULL,
				amount DECIMAL(10, 2) NOT NULL,
				method VARCHAR(255) NOT NULL,
				status VARCHAR(255) NOT NULL,
				created_at TIMESTAMP NOT NULL,
				updated_at TIMESTAMP NOT NULL
			);

			CREATE INDEX idx_payments_order_id ON payments (order_id);

			ALTER TABLE payments ADD CONSTRAINT fk_payments_order_id FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE;
		`

		if err := tx.Exec(statement).Error; err != nil {
			return fmt.Errorf("Error while creating payments table: %v", err)
		}

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Migrator().DropTable("payments"); err != nil {
			return fmt.Errorf("Error while dropping payments table: %v", err)
		}

		return nil
	},
}
