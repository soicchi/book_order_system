package migrations

import (
	"fmt"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var CreateShippingAddress = &gormigrate.Migration{
	ID: "20241113121208",
	Migrate: func(tx *gorm.DB) error {
		type shippingAddress struct {
			ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
			CustomerID uuid.UUID `gorm:"type:uuid;not null;index"`
			Prefecture string    `gorm:"type:varchar(255);not null"`
			City       string    `gorm:"type:varchar(255);not null"`
			State      string    `gorm:"type:varchar(255);not null"`
			CreatedAt  time.Time `gorm:"autoCreateTime;not null"`
			UpdatedAt  time.Time `gorm:"autoUpdateTime;not null"`
		}

		if err := tx.Migrator().CreateTable(&shippingAddress{}); err != nil {
			return fmt.Errorf("failed to create table: %w", err)
		}

		return tx.Exec("ALTER TABLE shipping_addresses ADD CONSTRAINT fk_customer_id FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE").Error
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("shipping_addresses")
	},
}
