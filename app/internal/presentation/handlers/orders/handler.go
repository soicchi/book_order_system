package orders

import (
	"log/slog"
	"net/http"

	"github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/logging"
	"github.com/soicchi/book_order_system/internal/presentation/validator"
	"github.com/soicchi/book_order_system/internal/usecase/orders"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UseCase interface {
	CreateOrder(ctx echo.Context, dto *orders.CreateInput) error
	CancelOrder(ctx echo.Context, orderID uuid.UUID) error
}

type OrderHandler struct {
	useCase UseCase
	logger  logging.Logger
}

func NewHandler(useCase UseCase, logger logging.Logger) *OrderHandler {
	return &OrderHandler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *OrderHandler) CreateOrder(ctx echo.Context) error {
	var req CreateRequest

	if err := validator.BindAndValidate(ctx, &req); err != nil {
		h.logger.Error("failed to bind and validate", slog.Any("error", err.Error()))
		return errors.ReturnJSON(ctx, err)
	}

	detailDTO := make([]*orders.CreateDetailInput, 0, len(req.OrderDetails))
	for _, d := range req.OrderDetails {
		detailDTO = append(detailDTO, orders.NewCreateDetailInput(d.BookID, d.Quantity, d.Price))
	}

	dto := orders.NewCreateInput(req.UserID, detailDTO)

	if err := h.useCase.CreateOrder(ctx, dto); err != nil {
		h.logger.Error("failed to create order", slog.Any("error", err.Error()))
		return errors.ReturnJSON(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, nil)
}

func (h *OrderHandler) CancelOrder(ctx echo.Context) error {
	var req CancelRequest

	if err := validator.BindAndValidate(ctx, &req); err != nil {
		h.logger.Error("failed to bind and validate", slog.Any("error", err.Error()))
		return errors.ReturnJSON(ctx, err)
	}

	if err := h.useCase.CancelOrder(ctx, req.OrderID); err != nil {
		h.logger.Error("failed to cancel order", slog.Any("error", err.Error()))
		return errors.ReturnJSON(ctx, err)
	}

	return ctx.JSON(http.StatusOK, nil)
}
