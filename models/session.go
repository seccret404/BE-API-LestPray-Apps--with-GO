package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model

	ID	uuid.UUID	`gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt	time.Time `json:"created_at"`
}