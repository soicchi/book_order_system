package book

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func NewMockRepository() *MockRepository {
	return &MockRepository{}
}

func (m *MockRepository) Create(ctx echo.Context, book *Book) error {
	args := m.Called(ctx, book)
	return args.Error(0)
}

func (m *MockRepository) FindByID(ctx echo.Context, id uuid.UUID) (*Book, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Book), args.Error(1)
}

func (m *MockRepository) FindAll(ctx echo.Context) ([]*Book, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*Book), args.Error(1)
}

func (m *MockRepository) FindByIDs(ctx echo.Context, ids []uuid.UUID) ([]*Book, error) {
	args := m.Called(ctx, ids)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*Book), args.Error(1)
}

func (m *MockRepository) Update(ctx echo.Context, book *Book) error {
	args := m.Called(ctx, book)
	return args.Error(0)
}

func (m *MockRepository) Delete(ctx echo.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
