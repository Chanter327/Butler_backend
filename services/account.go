package services

import (
	"fmt"

	"gorm.io/gorm"

	config "github.com/Chanter327/Butler_backend/config"
	models "github.com/Chanter327/Butler_backend/models"
)

var db *gorm.DB

func init() {
	var err error
	db, err = config.ConnectDB()
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	if err := db.AutoMigrate(&models.Users{}); err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}
}

func Authentication(email, password string) (models.Users, error) {
	var user models.Users

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Users{}, fmt.Errorf("user not found")
		}
		return models.Users{}, fmt.Errorf("database error: %v", err)
	}

	if !CheckPasswordHash(password, user.Password) {
		return models.Users{}, fmt.Errorf("wrong password")
	}

	return user, nil
}

type RegistrationRes struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	UserName string `json:"userName,omitempty"`
}

func RegisterUser(user models.Users) (RegistrationRes, error) {
	newUser := user

	if err := db.Create(&newUser).Error; err != nil {
		res := RegistrationRes{
			Status: "fail",
			Message: "registration failed",
		}
		return res, err
	}

	res := RegistrationRes{
		Status: "success",
		Message: "registered user successfully",
	}

	return res, nil
}