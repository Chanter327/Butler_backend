package services

import (
	"log"

	main "github.com/Chanter327/Butler_backend"
	"github.com/Chanter327/Butler_backend/models"
)

func ShowUser() models.Users {
	db, err := main.ConnectDB(".env")
	if err != nil {
		log.Printf("failed to connect to database: %v", err)
		return models.Users{}
	}

	db.AutoMigrate(&models.Users{})

	if err := db.AutoMigrate(&models.Users{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	var user models.Users
	db.First(&user, 1)

	if err := db.First(&user, 1).Error; err != nil {
		log.Printf("failed to find user: %v", err)
	}

	return user
}