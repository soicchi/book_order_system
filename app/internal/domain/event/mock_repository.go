package event

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockEventRepository struct {
	mock.Mock
}

func NewMockEventRepository() *MockEventRepository {
	return &MockEventRepository{}
}

func (m *MockEventRepository) Create(ctx echo.Context, event *Event) error {
	args := m.Called(ctx, event)
	return args.Error(0)
}

func (m *MockEventRepository) FetchByID(ctx echo.Context, eventID uuid.UUID) (*Event, error) {
	args := m.Called(ctx, eventID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Event), args.Error(1)
}

func (m *MockEventRepository) FetchAll(ctx echo.Context) ([]*Event, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*Event), args.Error(1)
}

func (m *MockEventRepository) Update(ctx echo.Context, event *Event) error {
	args := m.Called(ctx, event)
	return args.Error(0)
}
