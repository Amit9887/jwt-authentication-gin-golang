package initializers

import "github.com/Amit9887/go-jwt-authentication/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
