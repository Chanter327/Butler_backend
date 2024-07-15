package services

import (
	"fmt"
	"net/http"

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

type LoginRes struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	UserName string `json:"userName,omitempty"`
	Code int
}

func Authentication(email, password string) LoginRes {
	var user models.Users

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			res := LoginRes{
				Status: "fail",
				Message: "user not found",
				Code: http.StatusUnauthorized,
			}
			return res
		}
		res := LoginRes{
			Status: "fail",
			Message: "database error: " + err.Error(),
		}
		return res
	}

	if !CheckPasswordHash(password, user.Password) {
		res := LoginRes{
			Status: "fail",
			Message: "wrong password",
			Code: http.StatusUnauthorized,
		}
		return res
	}

	res := LoginRes{
		Status: "success",
		Message: "loged in successfully",
		UserName: user.UserName,
		Code: http.StatusOK,
	}
	return res
}

type RegistrationRes struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	UserName string `json:"userName,omitempty"`
	Code int
}

func RegisterUser(user models.Users) RegistrationRes {
    var newUser models.Users = user
    var existingUser models.Users

    // ユーザーがすでに登録されている確認
    if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
        if err != gorm.ErrRecordNotFound {
            res := RegistrationRes{
                Status: "fail",
                Message: fmt.Sprintf("database error: %v", err),
				Code: http.StatusInternalServerError,
            }
            return res
        }
    } else {
        res := RegistrationRes{
            Status: "fail",
            Message: "user already registered",
			Code: http.StatusConflict,
        }
        return res
    }

    if err := db.Create(&newUser).Error; err != nil {
        res := RegistrationRes{
            Status: "fail",
            Message: "registration failed",
			Code: http.StatusInternalServerError,
        }
        return res
    }

    res := RegistrationRes{
        Status: "success",
        Message: "registered user successfully",
		Code: http.StatusOK,
    }

    return res
}
