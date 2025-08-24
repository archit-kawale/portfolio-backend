package models

import "time"

type Message struct {
	ID        int64      `json:"id" gorm:"primary_key"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Message   string     `json:"message"`
	CreatedAt *time.Time `json:"created_at" gorm:"autoCreateTime"` // Time type should be pointer
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"` // Time type should be pointer
}
