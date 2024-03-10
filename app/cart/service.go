package cart

import (
	"github.com/lai0xn/cr-dermasuim/models"
	"github.com/lai0xn/cr-dermasuim/storage"
	uuid "github.com/satori/go.uuid"
)

type Service struct{}

func (s *Service) AddToCart(userID uuid.UUID, productID uuid.UUID) error {
	var product models.Product
	var user models.User
	db := storage.DB.Where("id = ?", userID).Find(&user)
	if db.Error != nil {
		return db.Error
	}
	db = storage.DB.Where("id = ?", productID).Find(&product)
	if db.Error != nil {
		return db.Error
	}
	if user.Cart.Items != nil {
		storage.DB.Create(&models.Cart{
			UserID: user.ID,
		})
	}
	item := &models.Item{
		ProductID: productID,
		Quantity:  0,
		UserID:    user.ID,
	}
	storage.DB.Create(&item)
	user.Cart.Items = append(user.Cart.Items, *item)
	storage.DB.Save(&user)
	return nil
}
