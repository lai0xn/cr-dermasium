package users

import (
	"github.com/lai0xn/cr-dermasuim/models"
	"github.com/lai0xn/cr-dermasuim/storage"
	uuid "github.com/satori/go.uuid"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetUserProfile(id uuid.UUID) (models.User, error) {
	var user models.User
	err := storage.DB.Where("id = ?").First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *Service) SetUserProfile(id uuid.UUID, payload ProfilePayload) error {
	var user models.User
	db := storage.DB.Where("id = ?", id).Find(&user)
	if db.Error != nil {
		return db.Error
	}
	user.FirstName = payload.FirstName
	user.LastName = payload.LastName
	user.Email = payload.Email
	user.Age = payload.Age
	user.Adress = payload.Adress
	db = storage.DB.Save(&user)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
