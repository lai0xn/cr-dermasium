package items

import (
	"github.com/lai0xn/cr-dermasuim/models"
	"github.com/lai0xn/cr-dermasuim/storage"
	uuid "github.com/satori/go.uuid"
)

type Service struct{}

func (s *Service) GetItemByID(id uuid.UUID) (*models.Item, error) {
	var item models.Item
	db := storage.DB.Where("id = ?", id).Find(&item)
	if db.Error != nil {
		return nil, db.Error
	}
	return &item, nil
}

func (s *Service) IncreaseQuantity(id uuid.UUID) error {
	var item models.Item
	db := storage.DB.Where("id = ?", id).Find(&item)
	if db.Error != nil {
		return db.Error
	}
	item.Quantity++
	storage.DB.Save(&item)
	return nil
}

func (s *Service) DecreaseQuantity(id uuid.UUID) error {
	var item models.Item
	db := storage.DB.Where("id = ?", id).Find(&item)
	if db.Error != nil {
		return db.Error
	}
	item.Quantity--
	db = storage.DB.Save(&item)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (s *Service) DeleteItem(id uuid.UUID) error {
	var item models.Item
	db := storage.DB.Where("id = ?", id).Find(&item)
	if db.Error != nil {
		return db.Error
	}
	db = storage.DB.Delete(&item)

	if db.Error != nil {
		return db.Error
	}
	return nil
}
