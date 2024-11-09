package database

import (
	"fmt"

	"github.com/soicchi/book_order_system/config"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	host     string
	name     string
	user     string
	password string
	port     string
	sslMode  string
}

func NewDBConfig(cfg *config.Config) *DBConfig {
	return &DBConfig{
		host:     cfg.DBHost,
		name:     cfg.DBName,
		user:     cfg.DBUser,
		password: cfg.DBPassword,
		port:     cfg.DBPort,
		sslMode:  cfg.DBSSLMode,
	}
}

var db *gorm.DB

func (d *DBConfig) Connect() error {
	dsn := d.dsn()
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		// this option is required to return the errors which are defined in the gorm.io/gorm package
		// https://github.com/go-gorm/gorm/blob/deceebfab8c460cfee229233aded2821ac6b08eb/errors.go#L10-L53
		TranslateError: true,
	})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	return nil
}

func (d *DBConfig) dsn() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		d.host, d.user, d.password, d.name, d.port, d.sslMode,
	)
}

type DBConnector struct{}

func NewDBConnector() *DBConnector {
	return &DBConnector{}
}

func (dc *DBConnector) GetDB(ctx echo.Context) *gorm.DB {
	// TODO: return tx if exists in context
	// TODO: we plan to implement this when we implement the transaction management
	return db
}
