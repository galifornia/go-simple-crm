package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name    string    `json:"name"`
	Company string    `json:"company"`
	Email   string    `json:"email"`
	Phone   int       `json:"phone"`
}
