package customers

import (
	"github.com/soicchi/book_order_system/internal/domain/entity"
	"github.com/soicchi/book_order_system/internal/domain/interfaces"
	"github.com/soicchi/book_order_system/internal/logging"
	"github.com/soicchi/book_order_system/internal/usecase/dto"

	"github.com/labstack/echo/v4"
)

type CustomerUseCase struct {
	repository interfaces.CustomerRepository
	logger     logging.Logger
}

func NewCustomerUseCase(repository interfaces.CustomerRepository, logger logging.Logger) *CustomerUseCase {
	return &CustomerUseCase{
		repository: repository,
		logger:     logger,
	}
}

func (uc *CustomerUseCase) Execute(ctx echo.Context, dto dto.CreateCustomerInput) error {
	customer, err := entity.NewCustomer(dto.Name, dto.Email, dto.Password)
	if err != nil {
		return err
	}

	if err := uc.repository.Create(ctx, customer); err != nil {
		return err
	}

	return nil
}
