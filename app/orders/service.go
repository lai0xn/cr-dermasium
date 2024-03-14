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
	// use raw sql query to delete all items
	if err := tx.Exec("DELETE FROM items WHERE cart_id = ?", user.Cart.ID).Error; err != nil {
		return models.Order{}, err
	}
	return order, tx.Commit().Error
}

func (s *Service) HandleWebhook(event string, orderID uuid.UUID) {
	var order models.Order
	fmt.Println("connected to webhook")
	fmt.Println(event)

	fmt.Println(orderID.String())
	if err := storage.DB.Where("id = ?", orderID).Find(&order).Error; err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	fmt.Println(event)
	switch event {
	case "checkout.paid":
		fmt.Println("the payment is successful")
		order.IsPayed = true
		storage.DB.Save(&order)
	}
}

func (s *Service) PayOrder() {
}
