package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterHRRoutes(router *gin.Engine) {
	hr := router.Group("/hr")
	RegisterVersionRoutes(hr)
}
