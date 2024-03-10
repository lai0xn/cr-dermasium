package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserID    uuid.UUID
	User      User
	ProductID uuid.UUID
	Product   Product
	CartID    uuid.UUID
	Quantity  int
}

func (u *Item) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.NewV4()
	return nil
}
