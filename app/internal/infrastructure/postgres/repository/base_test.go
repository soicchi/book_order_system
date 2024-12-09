package repository

import (
	"fmt"
	"log"
	"os"
	"testing"

	"event_system/internal/infrastructure/postgres/database"

	"github.com/labstack/echo/v4"
)

func TestMain(m *testing.M) {
	dbHost := os.Getenv("TEST_DB_HOST")
	dbUser := os.Getenv("TEST_DB_USER")
	dbPassword := os.Getenv("TEST_DB_PASSWORD")
	dbName := os.Getenv("TEST_DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPassword, dbName, "5432", "disable",
	)
	if err := database.SetupTestDB(dsn); err != nil {
		log.Fatalf("failed to setup test database: %v", err)
	}

	e := echo.New()
	ctx := e.NewContext(nil, nil)
	db := database.GetDB(ctx)

	defer func() {
		statement := "DROP SCHEMA public CASCADE; CREATE SCHEMA public;"

		if r := recover(); r != nil {
			if err := db.Exec(statement).Error; err != nil {
				log.Fatalf("failed to drop schema: %v", err)
			}
		}

		if err := db.Exec(statement).Error; err != nil {
			log.Fatalf("failed to drop schema: %v", err)
		}
	}()

	if err := database.Migrate(); err != nil {
		log.Fatalf("failed to migrate test database: %v", err)
	}

	if err := database.CreateTestData(); err != nil {
		log.Fatalf("failed to create test data: %v", err)
	}

	log.Println("Test database is ready")

	m.Run()
}
