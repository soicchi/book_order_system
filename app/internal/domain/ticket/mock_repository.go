package ticket

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

func (r *MockRepository) Create(ctx echo.Context, ticket *Ticket) error {
	args := r.Called(ctx, ticket)
	return args.Error(0)
}

func (r *MockRepository) FetchByRegistrationID(ctx echo.Context, registrationID uuid.UUID) (*Ticket, error) {
	args := r.Called(ctx, registrationID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Ticket), args.Error(1)
}

func (r *MockRepository) Update(ctx echo.Context, ticket *Ticket) error {
	args := r.Called(ctx, ticket)
	return args.Error(0)
}
