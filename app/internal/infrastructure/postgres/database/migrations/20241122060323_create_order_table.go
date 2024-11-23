package migrations

import (
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var CreateOrderTable = &gormigrate.Migration{
	ID: "20241122060323",
	Migrate: func(tx *gorm.DB) error {
		statement := `
			CREATE TABLE IF NOT EXISTS orders (
				id UUID PRIMARY KEY,
				customer_id UUID NOT NULL,
				total_price DECIMAL(10, 2) NOT NULL,
				ordered_at TIMESTAMP NOT NULL
			);

			CREATE INDEX idx_orders_customer_id ON orders (customer_id);
			CREATE INDEX idx_orders_ordered_at ON orders (ordered_at);

			ALTER TABLE orders ADD CONSTRAINT fk_orders_customer_id FOREIGN KEY (customer_id) REFERENCES customers (id) ON DELETE CASCADE;
		`

		if err := tx.Exec(statement).Error; err != nil {
			return fmt.Errorf("Error while creating orders table: %v", err)
		}

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Migrator().DropTable("orders"); err != nil {
			return fmt.Errorf("Error while dropping orders table: %v", err)
		}

		return nil
	},
}
