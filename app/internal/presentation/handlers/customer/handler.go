package customer

import (
	"log/slog"
	"net/http"

	"github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/logging"
	"github.com/soicchi/book_order_system/internal/usecase/customers"
	"github.com/soicchi/book_order_system/internal/usecase/dto"

	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	useCase *customers.CustomerUseCase
	logger  logging.Logger
}

func NewCustomerHandler(useCase *customers.CustomerUseCase, logger logging.Logger) *CustomerHandler {
	return &CustomerHandler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *CustomerHandler) CreateCustomer(ctx echo.Context) error {
	var req CreateCustomerRequest

	if err := ctx.Bind(&req); err != nil {
		h.logger.Error("failed to bind request", slog.String("error", err.Error()))
		return err.(*errors.CustomError).ReturnJSON(ctx)
	}

	if err := ctx.Validate(req); err != nil {
		h.logger.Error("validation error", slog.String("error", err.Error()))
		return err.(*errors.CustomError).ReturnJSON(ctx)
	}

	dto := dto.CreateCustomerInput{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.useCase.Execute(ctx, dto); err != nil {
		h.logger.Error("failed to create customer", slog.String("error", err.Error()))
		return err.(*errors.CustomError).ReturnJSON(ctx)
	}

	return ctx.JSON(http.StatusCreated, "created customer successfully")
}
