package user

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{}
}

func (m *MockUserRepository) Create(ctx echo.Context, user *User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) FetchByID(ctx echo.Context, userID uuid.UUID) (*User, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) FetchAll(ctx echo.Context) ([]*User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*User), args.Error(1)
}

func (m *MockUserRepository) Update(ctx echo.Context, user *User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}
