package models

import (
	"github.com/google/uuid"
)

// User model
type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Email    string    `json:"email" gorm:"required"`
	Password string    `json:"password" gorm:"required"`
	Todos    []Todo    `json:"todos" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
