package customers

import (
	"github.com/soicchi/book_order_system/internal/domain/interfaces"
	"github.com/soicchi/book_order_system/internal/logging"
	"github.com/soicchi/book_order_system/internal/usecase/dto"

	"github.com/labstack/echo/v4"
)

type FetchCustomerUseCase struct {
	repository interfaces.CustomerRepository
	logger     logging.Logger
}

func NewFetchCustomerUseCase(repository interfaces.CustomerRepository, logger logging.Logger) *FetchCustomerUseCase {
	return &FetchCustomerUseCase{
		repository: repository,
		logger:     logger,
	}
}

func (uc *FetchCustomerUseCase) Execute(ctx echo.Context, id string) (*dto.FetchCustomerOutput, error) {
	customer, err := uc.repository.FetchByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.FetchCustomerOutput{
		ID:        customer.ID().String(),
		Name:      customer.Name(),
		Email:     customer.Email(),
		CreatedAt: *customer.CreatedAt(),
		UpdatedAt: *customer.UpdatedAt(),
	}, nil
}
