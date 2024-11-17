package orders

import (
	"net/http"

	"github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/logging"
	"github.com/soicchi/book_order_system/internal/usecase/dto"
	"github.com/soicchi/book_order_system/internal/usecase/orders"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	useCase *orders.OrderUseCase
	logger  logging.Logger
}

func NewOrderHandler(useCase *orders.OrderUseCase, logger logging.Logger) *OrderHandler {
	return &OrderHandler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *OrderHandler) CreateOrder(ctx echo.Context) error {
	var req CreateOrderRequest

	if err := ctx.Bind(&req); err != nil {
		h.logger.Error("failed to bind request", "error", err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "failed to bind request"})
	}

	if err := ctx.Validate(req); err != nil {
		h.logger.Error("validation error", "error", err.Error())
		return err.(*errors.CustomError).ReturnJSON(ctx)
	}

	dto := &dto.CreateOrderInput{
		CustomerID:        req.CustomerID,
		ShippingAddressID: req.ShippingAddressID,
	}

	if err := h.useCase.CreateOrder(ctx, dto); err != nil {
		h.logger.Error("failed to create order", "error", err.Error())
		return err.(*errors.CustomError).ReturnJSON(ctx)
	}

	return ctx.JSON(http.StatusCreated, map[string]string{"message": "created order successfully"})
}
