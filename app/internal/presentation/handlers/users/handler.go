package users

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/logging"
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

func NewUserHandler(useCase UseCase, logger logging.Logger) *UserHandler {
	return &UserHandler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *UserHandler) CreateUser(ctx echo.Context) error {
	var req CreateRequest

	// Echo の Bind 関数はJSON形式や型のチェックを行うものでユーザーからの入力を検証するものではないので、
	// ValidationErrorとだけ返す。
	if err := ctx.Bind(&req); err != nil {
		h.logger.Error("failed to bind request", slog.Any("error", err))
		customErr := errors.New(fmt.Errorf("failed to bind request: %w", err), errors.ValidationError)
		return errors.ReturnJSON(ctx, customErr)
	}

	if err := ctx.Validate(&req); err != nil {
		h.logger.Error("failed to validate request", slog.Any("error", err))
		return errors.ReturnJSON(ctx, err)
	}

	dto := dto.NewCreateInput(req.Name, req.Email, req.Password)

	if err := h.useCase.CreateUser(ctx, dto); err != nil {
		h.logger.Error("failed to create user", slog.Any("error", err))
		return errors.ReturnJSON(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, nil)
}
