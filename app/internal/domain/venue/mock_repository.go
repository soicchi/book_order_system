package venue

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockVenueRepository struct {
	mock.Mock
}

func NewMockVenueRepository() *MockVenueRepository {
	return &MockVenueRepository{}
}

func (m *MockVenueRepository) Create(ctx echo.Context, v *Venue) error {
	args := m.Called(ctx, v)
	return args.Error(0)
}

func (m *MockVenueRepository) FetchAll(ctx echo.Context) ([]*Venue, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*Venue), args.Error(1)
}

func (m *MockVenueRepository) FetchByID(ctx echo.Context, venueID uuid.UUID) (*Venue, error) {
	args := m.Called(ctx, venueID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Venue), args.Error(1)
}

func (m *MockVenueRepository) Update(ctx echo.Context, v *Venue) error {
	args := m.Called(ctx, v)
	return args.Error(0)
}
