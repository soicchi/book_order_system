package user

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func NewRepository() *MockRepository {
	return &MockRepository{}
}

func (m *MockRepository) Create(ctx echo.Context, user *User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockRepository) FindByID(ctx echo.Context, id uuid.UUID) (*User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*User), args.Error(1)
}

func (m *MockRepository) FindAll(ctx echo.Context) ([]*User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*User), args.Error(1)
}

func (m *MockRepository) Update(ctx echo.Context, user *User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockRepository) Delete(ctx echo.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
