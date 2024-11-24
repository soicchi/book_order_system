package database

import (
	"fmt"
	"strconv"
	"time"

	"github.com/soicchi/book_order_system/internal/config"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/database/migrations"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	host         string
	name         string
	user         string
	password     string
	port         string
	sslMode      string
	maxConnPool  int
	poolLifetime time.Duration
}

func NewDBConfig(cfg *config.Config) *DBConfig {
	maxConnPool, err := strconv.Atoi(cfg.DBMAXConnPool)
	if err != nil {
		maxConnPool = 20
	}

	poolLifetime, err := strconv.Atoi(cfg.DBPoolLifetime)
	if err != nil {
		poolLifetime = 300
	}

	return &DBConfig{
		host:         cfg.DBHost,
		name:         cfg.DBName,
		user:         cfg.DBUser,
		password:     cfg.DBPassword,
		port:         cfg.DBPort,
		sslMode:      cfg.DBSSLMode,
		maxConnPool:  maxConnPool,
		poolLifetime: time.Duration(poolLifetime) * time.Second,
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

	if err := d.setupPool(); err != nil {
		return fmt.Errorf("failed to set up the connection pool: %w", err)
	}

	return nil
}

func (d *DBConfig) dsn() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		d.host, d.user, d.password, d.name, d.port, d.sslMode,
	)
}

func GetDB(ctx echo.Context) *gorm.DB {
	tx, ok := ctx.Get("tx").(*gorm.DB)
	if ok {
		return tx
	}

	return db
}

func BeginTx(ctx echo.Context) (*gorm.DB, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}

	ctx.Set("tx", tx)

	return tx, nil
}

func Migrate() error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// Add migrations here
		migrations.CreateUserTable,
		migrations.CreateBookTable,
		migrations.CreateOrderTable,
		migrations.CreateOrderDetailTable,
	})

	if err := m.Migrate(); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}

func (d *DBConfig) setupPool() error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database/sql DB: %w", err)
	}

	// Set the maximum number of pool connections
	sqlDB.SetMaxIdleConns(d.maxConnPool)
	// Set the recycle time of the pool connections
	sqlDB.SetConnMaxLifetime(d.poolLifetime * time.Second)

	return nil
}
