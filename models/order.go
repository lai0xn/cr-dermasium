package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;primary_key"`
	UserID     uuid.UUID
	CartID     uuid.UUID
	Cart       Cart
	User       User
	IsPayed    bool
	PayingTime time.Time
	TotalPrice int
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	o.ID = uuid.NewV4()
	return nil
}
