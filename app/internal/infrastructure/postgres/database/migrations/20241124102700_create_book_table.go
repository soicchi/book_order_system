package migrations

import (
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var CreateBookTable = &gormigrate.Migration{
	ID: "20241124102700",
	Migrate: func(tx *gorm.DB) error {
		statement := `
			CREATE TABLE IF NOT EXISTS books (
				id UUID PRIMARY KEY,
				title VARCHAR(255) NOT NULL,
				author VARCHAR(255) NOT NULL,
				price DECIMAL(10, 2) NOT NULL,
				stock INT NOT NULL,
				created_at TIMESTAMP NOT NULL,
				updated_at TIMESTAMP NOT NULL
			);

			CREATE INDEX idx_books_title ON books(title);
			CREATE INDEX idx_books_author ON books(author);
		`

		if err := tx.Exec(statement).Error; err != nil {
			return fmt.Errorf("failed to create books table: %w", err)
		}

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Migrator().DropTable("books"); err != nil {
			return fmt.Errorf("failed to drop books table: %w", err)
		}

		return nil
	},
}
