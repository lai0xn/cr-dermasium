package orders

import (
	"fmt"

	"github.com/lai0xn/cr-dermasuim/models"
	"github.com/lai0xn/cr-dermasuim/storage"
	uuid "github.com/satori/go.uuid"
)

type Service struct{}

func (s *Service) GetOrderByID(uuid.UUID) {
}

func (s *Service) GetAllOrders() {
}

func (s *Service) CheckOut(userID uuid.UUID) (models.Order, error) {
	var user models.User

	tx := storage.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := storage.DB.Where("id = ?", userID).Preload("Cart.Items.Product").First(&user).Error; err != nil {
		tx.Rollback()
		return models.Order{}, err
	}
	var price int = 0
	for _, item := range user.Cart.Items {
		price += item.Product.Price * item.Quantity
	}
	fmt.Println(price)
	order := models.Order{
		CartID:     user.Cart.ID,
		UserID:     user.ID,
		IsPayed:    false,
		TotalPrice: price,
	}
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return models.Order{}, err
	}

	return order, tx.Commit().Error
}

func (s *Service) HandleWebhook(payload WebhookPayload, orderID uuid.UUID) {
	var order models.Order
	fmt.Println("connected to webhook")
	tx := storage.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Where("id = ?", orderID).First(&order); err != nil {
		tx.Rollback()
		panic(err)
	}

	switch payload.Type {
	case "checkout.paid":
		order.IsPayed = true
		tx.Save(&order)
	}
}

func (s *Service) PayOrder() {
}
