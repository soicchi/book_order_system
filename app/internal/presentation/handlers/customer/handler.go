package customer

import (
	"log/slog"
	"net/http"

	"github.com/soicchi/book_order_system/internal/dto"
	"github.com/soicchi/book_order_system/internal/logging"
	"github.com/soicchi/book_order_system/internal/usecase/customers"

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
		h.logger.Error("failed to bind request", slog.Any("error", err))
		return ctx.JSON(http.StatusBadRequest, "failed to bind request")
	}

	if err := ctx.Validate(req); err != nil {
		h.logger.Error("validation error", slog.Any("error", err))
		return ctx.JSON(http.StatusBadRequest, "validation error")
	}

	dto := dto.CreateCustomerInput{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.useCase.Execute(ctx, dto); err != nil {
		h.logger.Error("failed to create customer", slog.Any("error", err))
		return ctx.JSON(http.StatusInternalServerError, "failed to create customer")
	}

	return ctx.JSON(http.StatusCreated, "created customer successfully")
}
