package items

import (
	"github.com/lai0xn/cr-dermasuim/app/products"
	"github.com/lai0xn/cr-dermasuim/app/users"
	uuid "github.com/satori/go.uuid"
)

type Item struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserID    uuid.UUID
	User      users.User
	ProductID uuid.UUID
	Product   products.Product
	Quantity  int
}
