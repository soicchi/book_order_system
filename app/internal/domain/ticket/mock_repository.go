package ticket

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

func (r *MockRepository) Create(ctx echo.Context, ticket *Ticket) error {
	args := r.Called(ctx, ticket)
	return args.Error(0)
}

func (r *MockRepository) FetchByQRCode(ctx echo.Context, qrCode string) (*Ticket, error) {
	args := r.Called(ctx, qrCode)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Ticket), args.Error(1)
}

func (r *MockRepository) Update(ctx echo.Context, ticket *Ticket) error {
	args := r.Called(ctx, ticket)
	return args.Error(0)
}
