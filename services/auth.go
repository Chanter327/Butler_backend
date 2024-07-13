package services

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	database "github.com/Chanter327/Butler_backend.git/models"
)

func ShowUser() database.Users {
	dsn := "host=localhost user=coffee password=pass dbname=butler_db port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

	db.AutoMigrate(&database.Users{})

	if err := db.AutoMigrate(&database.Users{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	var user database.Users
	db.First(&user, 1)

	if err := db.First(&user, 1).Error; err != nil {
		log.Printf("failed to find user: %v", err)
	}

	return user
}