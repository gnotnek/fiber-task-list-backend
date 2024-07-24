package models

import (
	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
}
