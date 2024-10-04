package health_check

import (
	"github.com/soicchi/book_order_system/presentation/responsehelper"

	"github.com/gin-gonic/gin"
)

func HealthCheck(ctx *gin.Context) {
	res := HealthCheckResponse{
		Message: "Health check OK",
	}

	responsehelper.ReturnStatusOK(ctx, res)
}
