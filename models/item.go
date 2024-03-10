package models

import (
	uuid "github.com/satori/go.uuid"
)

type Item struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserID    uuid.UUID
	User      User
	ProductID uuid.UUID
	Product   Product
	CartID    uuid.UUID
	Quantity  int
}
