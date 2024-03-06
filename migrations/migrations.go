package migrations

import (
	"log"

	"github.com/lai0xn/cr-dermasuim/app/users"
	"github.com/lai0xn/cr-dermasuim/storage"
)

func Migrate() {
	log.Println("Running Migrations .....")
	storage.DB.AutoMigrate(users.User{})
}
