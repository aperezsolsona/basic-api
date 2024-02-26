package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;"`
	Name      string    `json:"name"`
	PetTypeID uint      `json:"pet_type_id"`
	PetType   PetType   `gorm:"foreignKey:PetTypeID" json:"pet_type"`
}

type PetType struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
