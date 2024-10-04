package responsehelper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func ReturnStatusOK[T any](ctx *gin.Context, body T) {
	ctx.JSON(http.StatusOK, &body)
}

func ReturnStatusCreated[T any](ctx *gin.Context, body T) {
	ctx.JSON(http.StatusCreated, &body)
}

func ReturnStatusBadRequest[T any](ctx *gin.Context, body T) {
	ctx.JSON(http.StatusBadRequest, &body)
}

func ReturnStatusInternalServerError[T any](ctx *gin.Context, body T) {
	ctx.JSON(http.StatusInternalServerError, &body)
}
