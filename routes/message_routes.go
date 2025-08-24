package routes

import (
	"github.com/archit-kawale/portfolio-backend/controllers"

	"github.com/gin-gonic/gin"
)

// SetupMessageRoutes defines the POST route for messages
func SetupMessageRoutes(rg *gin.RouterGroup) {
	rg.POST("/message", controllers.HandlePostMessage)
}
