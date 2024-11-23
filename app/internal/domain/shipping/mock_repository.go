package shipping

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(ctx echo.Context, shipping *Shipping, orderID uuid.UUID) error {
	args := m.Called(ctx, shipping, orderID)
	return args.Error(0)
}

func (m *MockRepository) UpdateStatus(ctx echo.Context, shipping *Shipping) error {
	args := m.Called(ctx, shipping)
	return args.Error(0)
}
