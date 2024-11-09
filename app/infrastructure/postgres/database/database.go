package database

import (
	"fmt"

	"github.com/soicchi/book_order_system/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnector struct {
	host     string
	name     string
	user     string
	password string
	port     string
	sslMode  string
}

func NewDBConnector(cfg *config.Config) *DBConnector {
	return &DBConnector{
		host:     cfg.DBHost,
		name:     cfg.DBName,
		user:     cfg.DBUser,
		password: cfg.DBPassword,
		port:     cfg.DBPort,
		sslMode:  cfg.DBSSLMode,
	}
}

var db *gorm.DB

func (d *DBConnector) Connect() error {
	dsn := d.dsn()
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	return nil
}

func (d *DBConnector) dsn() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		d.host, d.user, d.password, d.name, d.port, d.sslMode,
	)
}
