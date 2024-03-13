package migrations

import (
	"log"

	"github.com/lai0xn/cr-dermasuim/models"
	"github.com/lai0xn/cr-dermasuim/storage"
)

func Migrate() {
	log.Println("Running Migrations .....")
	storage.DB.AutoMigrate(models.User{})
	storage.DB.AutoMigrate(models.Cart{})
	storage.DB.AutoMigrate(models.Item{})
	storage.DB.AutoMigrate(models.Product{})
	storage.DB.AutoMigrate(models.Order{})
}
