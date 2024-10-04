package customers

import (
	"log/slog"

	"github.com/soicchi/book_order_system/logger"
	"github.com/soicchi/book_order_system/presentation/responsehelper"
	customerUseCase "github.com/soicchi/book_order_system/usecase/customers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateCustomerHandler struct {
	useCase customerUseCase.CreateCustomerUseCase
	logger  logger.Logger
}

func NewCreateCustomerHandler(useCase customerUseCase.CreateCustomerUseCase, logger logger.Logger) *CreateCustomerHandler {
	return &CreateCustomerHandler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *CreateCustomerHandler) PostUser(ctx *gin.Context) {
	var req CreateCustomerRequest

	// bind request
	if err := ctx.ShouldBind(&req); err != nil {
		h.logger.Error("failed to bind request", slog.Any("error", err))
		res := responsehelper.ErrorResponse{
			Message: "failed to bind request customer",
		}
		responsehelper.ReturnStatusBadRequest(ctx, res)
	}

	// validate request
	validator := validator.New()
	if err := validator.Struct(&req); err != nil {
		h.logger.Error("failed to validate request", slog.Any("error", err))
		res := responsehelper.ErrorResponse{
			Message: "failed to validate request customer",
		}
		responsehelper.ReturnStatusBadRequest(ctx, res)
	}

	// create input DTO
	input := customerUseCase.CreateUseCaseInputDTO{
		Name:       req.Name,
		Email:      req.Email,
		Prefecture: req.Prefecture,
		Address:    req.Address,
		Password:   req.Password,
	}

	// execute use case
	if err := h.useCase.Execute(ctx, input); err != nil {
		h.logger.Error("failed to execute use case", slog.Any("error", err))
		res := responsehelper.ErrorResponse{
			Message: "failed to create customer",
		}
		responsehelper.ReturnStatusInternalServerError(ctx, res)
	}

	res := CreateCustomerResponse{
		Message: "success to create customer",
	}
	responsehelper.ReturnStatusCreated(ctx, res)
}
