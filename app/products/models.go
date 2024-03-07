package products

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name        string
	Description string
	Price       int
	Picture     string
	InStock     int
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.NewV4()
	return nil
}
