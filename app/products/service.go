package products

import (
	"errors"
	"fmt"

	"github.com/lai0xn/cr-dermasuim/models"
	"github.com/lai0xn/cr-dermasuim/storage"
	uuid "github.com/satori/go.uuid"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateProduct(p ProductPayload) error {
	db := storage.DB.Create(&models.Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		InStock:     p.InStock,
		Picture:     p.Picture,
	})
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (s *Service) DeleteProduct(id uuid.UUID) error {
	var product models.Product
	db := storage.DB.Where("id = ?", id).First(&product)
	if db.Error != nil {
		return db.Error
	}
	db = storage.DB.Delete(product)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (s *Service) UpdateProduct(id uuid.UUID, p ProductPayload) error {
	var product models.Product
	db := storage.DB.Where("id = ?", id).First(&product)

	if db.Error != nil {
		return db.Error
	}
	product.Price = p.Price
	product.Name = p.Name
	product.Description = p.Description
	product.InStock = p.InStock
	db = storage.DB.Save(&product)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (s *Service) GetProduct(id uuid.UUID) (models.Product, error) {
	var product models.Product
	db := storage.DB.Where("id = ?", id).First(&product)
	if db.Error != nil {
		return product, errors.New("product not found")
	}

	return product, nil
}

func (s *Service) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	db := storage.DB.Find(&products)
	if db.Error != nil {
		fmt.Println(db.Error)

		return nil, db.Error
	}
	return products, nil
}
