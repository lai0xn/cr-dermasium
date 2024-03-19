package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Message struct {
	text   string
	ChatID uuid.UUID
	Chat   Chat
}

type Chat struct {
	gorm.Model
	UserID   uuid.UUID
	DoctorID uuid.UUID
	Messages []Message
}
