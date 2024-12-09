package migrations

import (
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var AddEventTable = &gormigrate.Migration{
	ID: "20241204145847",
	Migrate: func(tx *gorm.DB) error {
		statement := `
			CREATE TABLE IF NOT EXISTS events (
				id UUID PRIMARY KEY,
				created_by UUID NOT NULL,
				venue_id UUID NOT NULL,
				title VARCHAR(255) NOT NULL,
				description TEXT NOT NULL,
				start_date TIMESTAMP NOT NULL,
				end_date TIMESTAMP NOT NULL,
				created_at TIMESTAMP NOT NULL,
				updated_at TIMESTAMP NOT NULL
			);

			ALTER TABLE events ADD CONSTRAINT fk_created_by FOREIGN KEY (created_by) REFERENCES users (id);
			ALTER TABLE events ADD CONSTRAINT fk_venue_id FOREIGN KEY (venue_id) REFERENCES venues (id);

			CREATE INDEX idx_events_created_by ON events (created_by);
		`

		if err := tx.Exec(statement).Error; err != nil {
			return fmt.Errorf("failed to create events table: %w", err)
		}

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Migrator().DropTable("events"); err != nil {
			return fmt.Errorf("failed to drop events table: %w", err)
		}

		return nil
	},
}
