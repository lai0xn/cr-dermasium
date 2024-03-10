package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ID     uuid.UUID
	UserID uuid.UUID
	User   *User
	Items  []Item
}
