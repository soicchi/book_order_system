package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DBConnector interface {
	GetDB(ctx echo.Context) *gorm.DB
}
