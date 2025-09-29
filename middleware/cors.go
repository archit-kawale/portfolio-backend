package middleware

import (
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	origins := os.Getenv("CORS_ORIGINS")
	allowOrigins := strings.Split(origins, ",")

	return cors.New(cors.Config{
		AllowOrigins: allowOrigins,
		AllowMethods: []string{"POST", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	})
}
