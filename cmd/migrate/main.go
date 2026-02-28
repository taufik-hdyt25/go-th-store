package migrate

import (
	"github.com/taufik-hdyt/go-crud/config"
	"github.com/taufik-hdyt/go-crud/models"
)

func main() {
	config.ConnectDatabase()

	config.DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
	)

	println("✅ Migration success")
}