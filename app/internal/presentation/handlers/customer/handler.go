package customer

import (
	"net/http"

	"github.com/soicchi/book_order_system/internal/dto"
	"github.com/soicchi/book_order_system/internal/logger"
	"github.com/soicchi/book_order_system/internal/usecase/customers"

	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	useCase *customers.CustomerUseCase
	logger  logger.Logger
}

func NewCustomerHandler(useCase *customers.CustomerUseCase, logger logger.Logger) *CustomerHandler {
	return &CustomerHandler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *CustomerHandler) CreateCustomer(ctx echo.Context) error {
	var req CreateCustomerRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	dto := dto.CreateCustomerInput{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.useCase.Execute(ctx, dto); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusCreated, nil)
}
