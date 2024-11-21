package shippingAddresses

import (
	"net/http"

	"github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/logging"
	"github.com/soicchi/book_order_system/internal/usecase/dto"
	"github.com/soicchi/book_order_system/internal/usecase/shippingAddresses"

	"github.com/labstack/echo/v4"
)

type ShippingAddressHandler struct {
	useCase *shippingAddresses.ShippingAddressUseCase
	logger  logging.Logger
}

func NewShippingAddressHandler(useCase *shippingAddresses.ShippingAddressUseCase, logger logging.Logger) *ShippingAddressHandler {
	return &ShippingAddressHandler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *ShippingAddressHandler) CreateShippingAddress(ctx echo.Context) error {
	var req CreateShippingAddressRequest

	if err := ctx.Bind(&req); err != nil {
		h.logger.Error("failed to bind request", "error", err.Error())
		return err.(*errors.CustomError).ReturnJSON(ctx)
	}

	if err := ctx.Validate(req); err != nil {
		h.logger.Error("validation error", "error", err.Error())
		return err.(*errors.CustomError).ReturnJSON(ctx)
	}

	dto, err := dto.NewCreateShippingAddressInput(req.CustomerID, req.Prefecture, req.City, req.State)
	if err != nil {
		h.logger.Error("failed to create dto", "error", err.Error())
		return err.(*errors.CustomError).ReturnJSON(ctx)
	}

	if err := h.useCase.CreateShippingAddress(ctx, dto); err != nil {
		h.logger.Error("failed to create shipping address", "error", err.Error())
		return err.(*errors.CustomError).ReturnJSON(ctx)
	}

	return ctx.JSON(http.StatusCreated, map[string]string{"message": "created shipping address successfully"})
}
