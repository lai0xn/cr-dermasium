package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Name     string
	Image    string
	Email    string
	Password string
}
