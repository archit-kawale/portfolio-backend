package main

import (
	"github.com/archit-kawale/portfolio-backend/config"
	"github.com/archit-kawale/portfolio-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	config.InitializeSnowflake(1)

	r := gin.Default()

	api := r.Group("/api")
	{
		routes.SetupMessageRoutes(api)
	}

	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
