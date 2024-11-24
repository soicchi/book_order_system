package migrations

import (
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var CreateOrderDetailTable = &gormigrate.Migration{
	ID: "20241124102333",
	Migrate: func(tx *gorm.DB) error {
		statement := `
			CREATE TABLE IF NOT EXISTS order_details (
				id UUID PRIMARY KEY,
				order_id UUID NOT NULL,
				book_id UUID NOT NULL,
				quantity INT NOT NULL,
				price DECIMAL(10, 2) NOT NULL
			);

			ALTER TABLE order_details ADD CONSTRAINT fk_order_id FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE;
			ALTER TABLE order_details ADD CONSTRAINT fk_book_id FOREIGN KEY (book_id) REFERENCES books(id);
		`

		if err := tx.Exec(statement).Error; err != nil {
			return fmt.Errorf("failed to create order_details table: %w", err)
		}

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Migrator().DropTable("order_details"); err != nil {
			return fmt.Errorf("failed to drop order_details table: %w", err)
		}
	
		return nil
	},
}
