package services

import (
	"github.com/archit-kawale/portfolio-backend/models"

	"github.com/archit-kawale/portfolio-backend/config"
)

func SaveMessage(name, email, message string) error {
	snowflakeID := config.Node.Generate().Int64()

	msg := models.Message{
		ID:      snowflakeID,
		Name:    name,
		Email:   email,
		Message: message,
	}
	if err := config.DB.Create(&msg).Error; err != nil {
		return err
	}
	return nil
}
