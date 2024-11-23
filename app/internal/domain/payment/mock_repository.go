package payment

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(ctx echo.Context, payment *Payment, orderID uuid.UUID) error {
	args := m.Called(ctx, payment, orderID)
	return args.Error(0)
}

func (m *MockRepository) UpdateStatus(ctx echo.Context, payment *Payment) error {
	args := m.Called(ctx, payment)
	return args.Error(0)
}
