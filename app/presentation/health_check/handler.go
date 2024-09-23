package health_check

import (
	"github.com/soicchi/book_order_system/presentation/response_helper"

	"github.com/gin-gonic/gin"
)

func HealthCheck(ctx *gin.Context) {
	res := HealthCheckResponse{
		Message: "Health check OK",
	}

	response_helper.ReturnStatusOK(ctx, res)
}
