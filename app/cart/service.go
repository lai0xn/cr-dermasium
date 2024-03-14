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
	var user models.User
	if err := storage.DB.Preload("Cart").Where("id = ?", userID).First(&user).Error; err != nil {
		return err
	}

	var product models.Product
	if err := storage.DB.Where("id = ?", productID).Find(&product).Error; err != nil {
		return err
	}

	if product.ID == uuid.Nil {
		return errors.New("product not found")
	}

	var cart *models.Cart
	if user.Cart == nil || user.Cart.ID == uuid.Nil {
		cart = &models.Cart{UserID: userID}
		if err := storage.DB.Create(cart).Error; err != nil {
			return err
		}
	} else {
		cart = user.Cart
	}

	var existingItem models.Item
	if err := storage.DB.Where("product_id = ? AND user_id = ? AND cart_id = ?", product.ID, userID, cart.ID).
		First(&existingItem).Error; err != nil &&
		!errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	tx := storage.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if existingItem.ID != uuid.Nil {
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

func (s *Service) DeleteCartItem(userID, itemID uuid.UUID) error {
	// Begin a transaction
	tx := storage.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Fetch the user and their cart
	var user models.User
	if err := tx.Preload("Cart").Where("id = ?", userID).First(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Check if the user has a cart
	if user.Cart == nil {
		tx.Rollback()
		return fmt.Errorf("user does not have a cart")
	}

	// Find the item in the cart
	var item models.Item
	if err := tx.Where("product_id = ? AND cart_id = ?", itemID, user.Cart.ID).First(&item).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("item not found in cart")
		}
		return err
	}

	// Delete the item
	if err := tx.Delete(&item).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) ViewCart(userID uuid.UUID) ([]models.Item, error) {
	var user models.User
	if err := storage.DB.Preload("Cart").Preload("Cart.Items").Preload("Cart.Items.Product").Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, err
		}
	}

	if user.Cart == nil {
		return nil, fmt.Errorf("user does not have a cart")
	}

	return user.Cart.Items, nil
}

func (s *Service) DecreaseItemQuantity(userID, itemID uuid.UUID) error {
	// Begin a transaction
	tx := storage.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Fetch the user and their cart
	var user models.User
	if err := tx.Preload("Cart").Preload("Cart.Items").Where("id = ?", userID).First(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Check if the user has a cart
	if user.Cart == nil {
		tx.Rollback()
		return fmt.Errorf("user does not have a cart")
	}

	// Find the item in the cart
	var item models.Item
	for _, cartItem := range user.Cart.Items {
		fmt.Println(cartItem.ID)
		if cartItem.ProductID == itemID {
			item = cartItem
			break
		}
	}
	if item.ID == uuid.Nil {
		tx.Rollback()
		return fmt.Errorf("item not found in cart")
	}

	// Decrease the quantity of the item
	if item.Quantity > 0 {
		item.Quantity--
		if err := tx.Save(&item).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
