package migrations

import (
	"fmt"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var CreateOrder = &gormigrate.Migration{
	ID: "20241117053721",
	Migrate: func(tx *gorm.DB) error {
		type order struct {
			ID                uuid.UUID `gorm:"type:uuid;primaryKey"`
			CustomerID        uuid.UUID `gorm:"type:uuid;not null;index"`
			ShippingAddressID uuid.UUID `gorm:"type:uuid;not null;index"`
			CreatedAt         time.Time `gorm:"autoCreateTime;not null;index"`
			UpdatedAt         time.Time `gorm:"autoUpdateTime;not null"`
		}

		if err := tx.Migrator().CreateTable(&order{}); err != nil {
			return fmt.Errorf("failed to create table: %w", err)
		}

		return tx.Exec("ALTER TABLE orders ADD CONSTRAINT fk_customer_id FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE SET NULL").Error
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("orders")
	},
}
