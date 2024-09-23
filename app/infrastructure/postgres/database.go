package postgres

import (
	"fmt"
	"os"

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
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		// TODO: I'm changing how to get the environment variables
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
}
