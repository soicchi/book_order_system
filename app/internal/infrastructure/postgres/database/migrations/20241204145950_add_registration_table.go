package migrations

import (
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var AddRegistrationTable = &gormigrate.Migration{
	ID: "20241204145950",
	Migrate: func(tx *gorm.DB) error {
		statement := `
			CREATE TABLE IF NOT EXISTS registrations (
				id UUID PRIMARY KEY,
				user_id UUID NOT NULL,
				event_id UUID NOT NULL,
				status VARCHAR(255) NOT NULL,
				registered_at TIMESTAMP NOT NULL
			);

			ALTER TABLE registrations ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id);
			ALTER TABLE registrations ADD CONSTRAINT fk_event_id FOREIGN KEY (event_id) REFERENCES events (id);

			CREATE INDEX idx_registrations_user_id ON registrations (user_id);
			CREATE INDEX idx_registrations_event_id ON registrations (event_id);
		`
		if err := tx.Exec(statement).Error; err != nil {
			return fmt.Errorf("failed to create registrations table: %w", err)
		}

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Migrator().DropTable("registrations"); err != nil {
			return fmt.Errorf("failed to drop registrations table: %w", err)
		}

		return nil
	},
}
