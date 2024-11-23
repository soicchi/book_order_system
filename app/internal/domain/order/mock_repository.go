package order

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(ctx echo.Context, order *Order, customerID uuid.UUID) error {
	args := m.Called(ctx, order, customerID)
	return args.Error(0)
}
