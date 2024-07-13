package services

import (
	"log"

	config "github.com/Chanter327/Butler_backend/config"
	models "github.com/Chanter327/Butler_backend/models"
)

func ShowUser() models.Users {
    db, err := config.ConnectDB()
    if err != nil {
        log.Printf("failed to connect to database: %v", err)
        return models.Users{}
    }

    // AutoMigrateを一度だけ呼び出す
    if err := db.AutoMigrate(&models.Users{}); err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }

    var user models.Users
    if err := db.First(&user, 1).Error; err != nil {
        log.Printf("failed to find user: %v", err)
        return models.Users{}  // ユーザーが見つからなかった場合は空の構造体を返す
    }

    return user
}
