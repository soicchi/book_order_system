package response_helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnStatusOK[T any](ctx *gin.Context, body T) {
	ctx.JSON(http.StatusOK, &body)
}
