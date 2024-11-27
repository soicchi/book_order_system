package books

import (
	"log/slog"
	"net/http"

	"github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/logging"
	"github.com/soicchi/book_order_system/internal/presentation/validator"
	"github.com/soicchi/book_order_system/internal/usecase/books"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UseCase interface {
	CreateBook(ctx echo.Context, dto *books.CreateInput) error
	RetrieveBook(ctx echo.Context, id uuid.UUID) (*books.RetrieveOutput, error)
	ListBooks(ctx echo.Context) ([]*books.ListOutput, error)
	UpdateBook(ctx echo.Context, dto *books.UpdateInput) error
}

type BookHandler struct {
	useCase UseCase
	logger  logging.Logger
}

func NewHandler(useCase UseCase, logger logging.Logger) *BookHandler {
	return &BookHandler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *BookHandler) CreateBook(ctx echo.Context) error {
	var req CreateRequest

	if err := validator.BindAndValidate(ctx, &req); err != nil {
		h.logger.Error("failed to bind and validate request", slog.Any("error", err))
		return errors.ReturnJSON(ctx, err)
	}

	dto := books.NewCreateInput(req.Title, req.Author, req.Price, req.Stock)

	if err := h.useCase.CreateBook(ctx, dto); err != nil {
		h.logger.Error("failed to create book", slog.Any("error", err))
		return errors.ReturnJSON(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, nil)
}

func (h *BookHandler) RetrieveBook(ctx echo.Context) error {
	var req RetrieveRequest

	if err := validator.BindAndValidate(ctx, &req); err != nil {
		h.logger.Error("failed to bind and validate request", slog.Any("error", err))
		return errors.ReturnJSON(ctx, err)
	}

	dto, err := h.useCase.RetrieveBook(ctx, req.ID)
	if err != nil {
		h.logger.Error("failed to retrieve book", slog.Any("error", err))
		return errors.ReturnJSON(ctx, err)
	}

	res := NewBookResponse(dto)

	return ctx.JSON(http.StatusOK, res)
}

func (h *BookHandler) ListBooks(ctx echo.Context) error {
	dto, err := h.useCase.ListBooks(ctx)
	if err != nil {
		h.logger.Error("failed to list books", slog.Any("error", err))
		return errors.ReturnJSON(ctx, err)
	}

	res := NewBooksResponse(dto)

	return ctx.JSON(http.StatusOK, res)
}

func (h *BookHandler) UpdateBook(ctx echo.Context) error {
	var req UpdateRequest

	if err := validator.BindAndValidate(ctx, &req); err != nil {
		h.logger.Error("failed to bind and validate request", slog.Any("error", err))
		return errors.ReturnJSON(ctx, err)
	}

	dto := books.NewUpdateInput(req.ID, req.Title, req.Author, req.Price, req.Stock)

	if err := h.useCase.UpdateBook(ctx, dto); err != nil {
		h.logger.Error("failed to update book", slog.Any("error", err))
		return errors.ReturnJSON(ctx, err)
	}

	return ctx.JSON(http.StatusOK, nil)
}
