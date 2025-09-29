package config

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB
var Node *snowflake.Node

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := os.Getenv("GO_ENV")
	if env == "prod" {
		gin.SetMode(gin.ReleaseMode) // Set Gin to release mode (no debug logs)
	} else {
		gin.SetMode(gin.DebugMode) // Set Gin to debug mode (verbose logging)
	}

	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	var errDb error
	DB, errDb = gorm.Open("postgres", dbURL)
	if errDb != nil {
		log.Fatal("Failed to connect to database:", errDb)
	}

	fmt.Println("Database connected successfully")
}

func InitializeSnowflake(nodeID int64) {
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		log.Fatalf("Error initializing snowflake node: %v", err)
	}
	Node = node
}
