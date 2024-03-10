package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primary_key;type:uuid"`
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	IsActive    bool
	Adress      string
	Cart        *Cart
	Age         int
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.NewV4()
	return nil
}
