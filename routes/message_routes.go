package routes

import (
	"github.com/archit-kawale/portfolio-backend/controllers"

	"github.com/gin-gonic/gin"
)

// SetupMessageRoutes defines the POST route for messages
func SetupMessageRoutes(r *gin.Engine) {
	r.POST("/message", controllers.HandlePostMessage)
}
