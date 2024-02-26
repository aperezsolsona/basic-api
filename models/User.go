package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}
