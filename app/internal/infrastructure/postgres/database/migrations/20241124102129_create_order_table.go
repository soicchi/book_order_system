package migrations

import (
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var CreateOrderTable = &gormigrate.Migration{
	ID: "20241124102129",
	Migrate: func(tx *gorm.DB) error {
		statement := `
			CREATE TABLE IF NOT EXISTS orders (
				id UUID PRIMARY KEY,
				user_id UUID NOT NULL,
				total_price DECIMAL(10, 2) NOT NULL,
				ordered_at TIMESTAMP NOT NULL
			);

			ALTER TABLE orders ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
		`

		if err := tx.Exec(statement).Error; err != nil {
			return fmt.Errorf("failed to create orders table: %w", err)
		}

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Migrator().DropTable("orders"); err != nil {
			return fmt.Errorf("failed to drop orders table: %w", err)
		}

		return nil
	},
}
