package postgres

import (
	"fmt"

	"github.com/soicchi/book_order_system/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Initialize(cfg config.Config) {
	dbURI := getDBURI(
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)
	db = sqlx.MustConnect("pgx", dbURI)
}

func GetDB() *sqlx.DB {
	return db
}

func getDBURI(dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode string) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode,
	)
}