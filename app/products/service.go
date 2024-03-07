package products

import (
	"github.com/lai0xn/cr-dermasuim/storage"
	uuid "github.com/satori/go.uuid"
)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) CreateProduct(p ProductPayload) error {
	db := storage.DB.Create(&Product{
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

func (s *service) DeleteProduct(id uuid.UUID) {
	var product Product
	storage.DB.Where("id = ?", id).First(&product)
	storage.DB.Delete(product)
}

func (s *service) UpdateProduct(id uuid.UUID, p ProductPayload) error {
	var product Product
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

func (s *service) GetProduct(id uuid.UUID) (Product, error) {
	var product Product
	db := storage.DB.Where("id = ?", id).First(&product)
	if db.Error != nil {
		return product, db.Error
	}

	return product, nil
}

func (s *service) GetAllProducts() ([]Product, error) {
	var products []Product
	db := storage.DB.First(&products)
	if db.Error != nil {
	}
	return nil, nil
}
