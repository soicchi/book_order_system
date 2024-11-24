package customers

import (
	"log/slog"
	"net/http"

	"github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/logging"
	"github.com/soicchi/book_order_system/internal/usecase/customers"

	"github.com/labstack/echo/v4"
)

type UseCase interface {
	CreateCustomer(ctx echo.Context, input *customers.CreateInput) error
}

type CustomerHandler struct {
	useCase UseCase
	logger  logging.Logger
}

func NewCustomerHandler(useCase UseCase, logger logging.Logger) *CustomerHandler {
	return &CustomerHandler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *CustomerHandler) CreateCustomer(ctx echo.Context) error {
	var input CreateCustomerRequest
	if err := ctx.Bind(&input); err != nil {
		h.logger.Error("failed to bind request body", slog.Any("error", err.Error()))
		return err.(*errors.CustomError).ReturnJSON(ctx)
	}

	if err := ctx.Validate(input); err != nil {
		h.logger.Error("failed to validate request body", slog.Any("error", err.Error()))
		return err.(*errors.CustomError).ReturnJSON(ctx)
	}

	dto := customers.NewCreateInput(input.Name, input.Email)

	if err := h.useCase.CreateCustomer(ctx, dto); err != nil {
		h.logger.Error("failed to create customer", slog.Any("error", err.Error()))
		return err
	}

	return ctx.JSON(http.StatusCreated, "customer created successfully")
}
