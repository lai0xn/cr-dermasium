package cart

import (
	"errors"
	"fmt"

	"github.com/lai0xn/cr-dermasuim/models"
	"github.com/lai0xn/cr-dermasuim/storage"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Service struct{}

func (s *Service) AddToCart(userID, productID uuid.UUID) error {
	tx := storage.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var user models.User
	if err := tx.Preload("Cart").Where("id = ?", userID).First(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	var cart *models.Cart
	if user.Cart == nil || user.Cart.ID == uuid.Nil {
		cart = &models.Cart{UserID: userID}
		if err := tx.Create(cart).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		cart = user.Cart
	}

	var existingItem models.Item
	if err := tx.Where("product_id = ? AND user_id = ? AND cart_id = ?",
		productID,
		userID,
		cart.ID).
		First(&existingItem).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return err
		}
	}

	if existingItem.ID != uuid.Nil {
		fmt.Println(existingItem)
		existingItem.Quantity++
		if err := tx.Save(&existingItem).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		newItem := &models.Item{
			ProductID: productID,
			Quantity:  1,
			UserID:    userID,
			CartID:    cart.ID,
		}

		if err := tx.Create(newItem).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
