package products

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       int
	Picture     string
	Rating      float32
	InStock     int
}
