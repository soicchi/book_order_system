package migrations

import (
	"fmt"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var CreateCartTable = &gormigrate.Migration{
	ID: "20241117143557",
	Migrate: func(tx *gorm.DB) error {
		type cart struct {
			ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
			Status     string    `gorm:"type:varchar(255);not null"`
			CreatedAt  time.Time `gorm:"autoCreateTime;not null"`
			UpdatedAt  time.Time `gorm:"autoUpdateTime;not null"`
			CustomerID uuid.UUID `gorm:"type:uuid;not null;index"`
		}

		if err := tx.Migrator().CreateTable(&cart{}); err != nil {
			return fmt.Errorf("failed to create table: %w", err)
		}

		return tx.Exec("ALTER TABLE carts ADD CONSTRAINT fk_customer_id FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE").Error
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("carts")
	},
}
