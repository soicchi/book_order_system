package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine) {
	basePath := r.Group("/api")
	v1 := basePath.Group("/v1")

	// /api/v1/health_check
	v1.GET("/health_check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Health check OK",
		})
	})
}
