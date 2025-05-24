package routes

import (
	"github.com/debadyuti0510/use-erp/internal/hr/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterVersionRoutes(rg *gin.RouterGroup) {
	rg.GET("/version", handlers.GetVersion)
}
