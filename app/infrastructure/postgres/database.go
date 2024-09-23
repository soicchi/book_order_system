package postgres

import (
	"fmt"

	"github.com/soicchi/book_order_system/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Initialize() {
	dbURI := getDBURI()
	db = sqlx.MustConnect("pgx", dbURI)
}

func GetDB() *sqlx.DB {
	return db
}

func getDBURI() string {
	cfg := config.GetConfig()

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)
}
