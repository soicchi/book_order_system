package users

import (
	"log/slog"
	"net/http"

	"github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/logging"
	"github.com/soicchi/book_order_system/internal/presentation/validator"
	dto "github.com/soicchi/book_order_system/internal/usecase/users"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UseCase interface {
	CreateUser(ctx echo.Context, dto *dto.CreateInput) error
	RetrieveUser(ctx echo.Context, id uuid.UUID) (*dto.RetrieveOutput, error)
	ListUsers(ctx echo.Context) ([]*dto.ListOutput, error)
	UpdateUser(ctx echo.Context, dto *dto.UpdateInput) error
	DeleteUser(ctx echo.Context, id uuid.UUID) error
}

type UserHandler struct {
	useCase UseCase
	logger  logging.Logger
}

func NewHandler(useCase UseCase, logger logging.Logger) *UserHandler {
	return &UserHandler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *UserHandler) CreateUser(ctx echo.Context) error {
	var req CreateRequest

	if err := validator.BindAndValidate(ctx, &req); err != nil {
		h.logger.Error("failed to bind and validate request", slog.Any("error", err))
		return errors.ReturnJSON(ctx, err)
	}

	dto := dto.NewCreateInput(req.Name, req.Email, req.Password)

	if err := h.useCase.CreateUser(ctx, dto); err != nil {
		h.logger.Error("failed to create user", slog.Any("error", err))
		return errors.ReturnJSON(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, nil)
}

func (h *UserHandler) RetrieveUser(ctx echo.Context) error {
	var req RetrieveRequest

	if err := validator.BindAndValidate(ctx, &req); err != nil {
		h.logger.Error("failed to bind and validate request", slog.Any("error", err))
		return errors.ReturnJSON(ctx, err)
	}

	dto, err := h.useCase.RetrieveUser(ctx, req.ID)
	if err != nil {
		h.logger.Error("failed to retrieve user", slog.Any("error", err))
		return errors.ReturnJSON(ctx, err)
	}

	res := NewUserResponse(dto)

	return ctx.JSON(http.StatusOK, res)
}

func (h *UserHandler) ListUsers(ctx echo.Context) error {
	dto, err := h.useCase.ListUsers(ctx)
	if err != nil {
		h.logger.Error("failed to list users", slog.Any("error", err))
		return errors.ReturnJSON(ctx, err)
	}

	res := NewUsersResponse(dto)

	return ctx.JSON(http.StatusOK, res)
}

func (h *UserHandler) UpdateUser(ctx echo.Context) error {
	var req UpdateRequest

	if err := validator.BindAndValidate(ctx, &req); err != nil {
		h.logger.Error("failed to bind and validate request", slog.Any("error", err))
		return errors.ReturnJSON(ctx, err)
	}

	dto := dto.NewUpdateInput(req.ID, req.Name, req.Email, req.Password)

	if err := h.useCase.UpdateUser(ctx, dto); err != nil {
		h.logger.Error("failed to update user", slog.Any("error", err))
		return errors.ReturnJSON(ctx, err)
	}

	return ctx.JSON(http.StatusOK, nil)
}

func (h *UserHandler) DeleteUser(ctx echo.Context) error {
	var req DeleteRequest

	if err := validator.BindAndValidate(ctx, &req); err != nil {
		h.logger.Error("failed to bind and validate request", slog.Any("error", err))
		return errors.ReturnJSON(ctx, err)
	}

	if err := h.useCase.DeleteUser(ctx, req.ID); err != nil {
		h.logger.Error("failed to delete user", slog.Any("error", err))
		return errors.ReturnJSON(ctx, err)
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
