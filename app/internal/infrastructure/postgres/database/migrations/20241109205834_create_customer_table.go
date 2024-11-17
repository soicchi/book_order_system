package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var CreateCustomerTable = &gormigrate.Migration{
	ID: "20241109205834",
	Migrate: func(tx *gorm.DB) error {
		type customer struct {
			ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
			Name      string    `gorm:"type:varchar(255);not null"`
			Email     string    `gorm:"type:varchar(255);not null;unique"`
			Password  string    `gorm:"type:varchar(255);not null"`
			CreatedAt time.Time `gorm:"autoCreateTime;not null"`
			UpdatedAt time.Time `gorm:"autoUpdateTime;not null"`
		}

		return tx.Migrator().CreateTable(&customer{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("customers")
	},
}
