package customer

import (
	"fmt"
	"log/slog"

	customerDomain "github.com/soicchi/book_order_system/domain/customer"

	"github.com/gin-gonic/gin"
)

type CreateCustomerUseCase struct {
	repository customerDomain.CustomerRepository
	logger     *slog.Logger
}

func NewCreateCustomerUseCase(repo customerDomain.CustomerRepository, logger *slog.Logger) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{
		repository: repo,
		logger:     logger,
	}
}

type CreateUseCaseInputDTO struct {
	Name       string
	Email      string
	Prefecture string
	Address    string
	Password   string
}

func (u *CreateCustomerUseCase) Execute(ctx *gin.Context, dto CreateUseCaseInputDTO) error {
	customer, err := customerDomain.NewCustomer(dto.Name, dto.Email, dto.Prefecture, dto.Address, dto.Password)
	if err != nil {
		return fmt.Errorf("failed to create customer domain: %w", err)
	}

	if err := u.repository.Create(customer); err != nil {
		return fmt.Errorf("failed to create customer in database: %w", err)
	}

	return nil
}
