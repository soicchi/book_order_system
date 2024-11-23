package customer

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func NewMockRepository() *MockRepository {
	return &MockRepository{}
}

func (m *MockRepository) Create(ctx echo.Context, customer *Customer) error {
	args := m.Called(ctx, customer)
	return args.Error(0)
}
