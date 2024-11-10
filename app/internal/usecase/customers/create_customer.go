package customers

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/entity"
	"github.com/soicchi/book_order_system/internal/domain/interfaces"
	"github.com/soicchi/book_order_system/internal/dto"
	"github.com/soicchi/book_order_system/internal/logger"

	"github.com/labstack/echo/v4"
)

type CustomerUseCase struct {
	repository interfaces.CustomerRepository
	logger     *logger.Logger
}

func NewCustomerUseCase(repository interfaces.CustomerRepository) *CustomerUseCase {
	return &CustomerUseCase{
		repository: repository,
	}
}

func (uc *CustomerUseCase) Execute(ctx echo.Context, dto dto.CreateCustomerInput) error {
	customer, err := entity.NewCustomer(dto.Name, dto.Email, dto.Password)
	if err != nil {
		return fmt.Errorf("failed to create customer entity: %w", err)
	}

	if err := uc.repository.Create(ctx, customer); err != nil {
		return err
	}

	return nil
}
