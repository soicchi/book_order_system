package postgres

import (
	"fmt"

	"github.com/soicchi/book_order_system/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Initialize(cfg config.Config) {
	dsn := getDSN(
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{}) 
}

func GetDB() *sqlx.DB {
	return db
}

func getDSN(dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode string) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode,
	)
}
