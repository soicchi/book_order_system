package router

import (
	"github.com/soicchi/book_order_system/config"
	"github.com/soicchi/book_order_system/presentation/health_check"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, cfg config.Config) {
	basePath := r.Group("/api")
	v1 := basePath.Group("/v1")

	// /api/v1/health_check
	v1.GET("/health_check", health_check.HealthCheck)
}
