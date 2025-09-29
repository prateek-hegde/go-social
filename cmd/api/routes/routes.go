package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateRoutes(server *gin.Engine) {
	v1Routes := server.Group("/v1")

	v1Routes.GET("/health", healthCheck)
}

func healthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "ok"})
}
