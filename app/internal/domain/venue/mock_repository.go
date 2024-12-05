package venue

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

func (m *MockRepository) Create(ctx echo.Context, v *Venue) error {
	args := m.Called(ctx, v)
	return args.Error(0)
}

func (m *MockRepository) FetchAll(ctx echo.Context) ([]*Venue, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*Venue), args.Error(1)
}

func (m *MockRepository) FetchByID(ctx echo.Context, venueID uuid.UUID) (*Venue, error) {
	args := m.Called(ctx, venueID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Venue), args.Error(1)
}

func (m *MockRepository) Update(ctx echo.Context, v *Venue) error {
	args := m.Called(ctx, v)
	return args.Error(0)
}
