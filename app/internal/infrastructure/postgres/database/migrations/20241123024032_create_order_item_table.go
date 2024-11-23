package migrations

import (
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var CreateOrderItemTable = &gormigrate.Migration{
	ID: "20241123024032",
	Migrate: func(tx *gorm.DB) error {
		statement := `
			CREATE TABLE IF NOT EXISTS order_items (
				id UUID PRIMARY KEY,
				order_id UUID NOT NULL,
				product_id UUID NOT NULL,
				quantity INT NOT NULL
			);

			CREATE INDEX idx_order_items_order_id ON order_items (order_id);
			CREATE INDEX idx_order_items_product_id ON order_items (product_id);

			ALTER TABLE order_items ADD CONSTRAINT fk_order_items_order_id FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE;
		`

		if err := tx.Exec(statement).Error; err != nil {
			return fmt.Errorf("Error while creating order_items table: %v", err)
		}

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Migrator().DropTable("order_items"); err != nil {
			return fmt.Errorf("Error while dropping order_items table: %v", err)
		}

		return nil
	},
}
