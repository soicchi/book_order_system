package registration

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockRegistrationRepository struct {
	mock.Mock
}

func NewMockRegistrationRepository() *MockRegistrationRepository {
	return &MockRegistrationRepository{}
}

func (m *MockRegistrationRepository) Create(ctx echo.Context, registration *Registration) error {
	args := m.Called(ctx, registration)
	return args.Error(0)
}

func (m *MockRegistrationRepository) Update(ctx echo.Context, registration *Registration) error {
	args := m.Called(ctx, registration)
	return args.Error(0)
}

func (m *MockRegistrationRepository) FetchByEventID(ctx echo.Context, eventID uuid.UUID) ([]*Registration, error) {
	args := m.Called(ctx, eventID)
	return args.Get(0).([]*Registration), args.Error(1)
}
