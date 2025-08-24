package controllers

import (
	"net/http"

	"github.com/archit-kawale/portfolio-backend/services"

	"github.com/gin-gonic/gin"
)

type MessageInput struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Message string `json:"message" binding:"required"`
}

// HandlePostMessage processes the POST request to save a message
func HandlePostMessage(c *gin.Context) {
	var input MessageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service layer to save message
	if err := services.SaveMessage(input.Name, input.Email, input.Message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Message saved successfully"})
}
