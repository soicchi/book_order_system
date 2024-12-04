package migrations

import (
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var AddTicketTable = &gormigrate.Migration{
	ID: "20241204145955",
	Migrate: func(tx *gorm.DB) error {
		statement := `
			CREATE TABLE IF NOT EXISTS tickets (
				id UUID PRIMARY KEY,
				registration_id UUID NOT NULL,
				qr_code VARCHAR(255) NOT NULL,
				status VARCHAR(255) NOT NULL,
				issued_at TIMESTAMP NOT NULL,
				used_at TIMESTAMP
			);

			ALTER TABLE tickets ADD CONSTRAINT fk_registration_id FOREIGN KEY (registration_id) REFERENCES registrations (id);

			CREATE INDEX idx_tickets_registration_id ON tickets (registration_id);
		`

		if err := tx.Exec(statement).Error; err != nil {
			return fmt.Errorf("failed to create tickets table: %w", err)
		}

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Migrator().DropTable("tickets"); err != nil {
			return fmt.Errorf("failed to drop tickets table: %w", err)
		}

		return nil
	},
}
