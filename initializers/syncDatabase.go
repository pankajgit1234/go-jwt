package initializers

import "example.com/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
