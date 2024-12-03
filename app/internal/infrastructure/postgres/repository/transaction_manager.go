package repository

import (
	"fmt"

	"event_system/internal/errors"
	"event_system/internal/infrastructure/postgres/database"

	"github.com/labstack/echo/v4"
)

type TransactionManager struct{}

func NewTransactionManager() *TransactionManager {
	return &TransactionManager{}
}

func (tm *TransactionManager) WithTransaction(ctx echo.Context, fn func(echo.Context) error) error {
	tx, err := database.BeginTx(ctx)
	if err != nil {
		return errors.New(
			fmt.Errorf("failed to begin transaction: %w", err),
			errors.UnexpectedError,
		)
	}

	if err := fn(ctx); err != nil {
		if err := tx.Rollback().Error; err != nil {
			return errors.New(
				fmt.Errorf("failed to rollback transaction: %w", err),
				errors.UnexpectedError,
			)
		}

		return nil
	}

	if err := tx.Commit().Error; err != nil {
		return errors.New(
			fmt.Errorf("failed to commit transaction: %w", err),
			errors.UnexpectedError,
		)
	}

	return nil
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
