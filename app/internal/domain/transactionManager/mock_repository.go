package transactionManager

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) WithTransaction(ctx echo.Context, fn func(ctx echo.Context) error) error {
	args := m.Called(ctx, fn)
	return args.Error(0)
}
