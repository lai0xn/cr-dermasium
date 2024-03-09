package cart

import (
	"github.com/lai0xn/cr-dermasuim/app/users"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ID     uuid.UUID
	UserID uuid.UUID
	User   users.User
}
