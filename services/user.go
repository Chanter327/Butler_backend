package services

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"

	config "github.com/Chanter327/Butler_backend/config"
	models "github.com/Chanter327/Butler_backend/models"
	structs "github.com/Chanter327/Butler_backend/structs"
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

func Authentication(email, password string) structs.LoginRes {
	var user models.Users

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			res := structs.LoginRes{
				Status: "fail",
				Message: "user not found",
				Code: http.StatusNotFound,
			}
			return res
		}
		res := structs.LoginRes{
			Status: "fail",
			Message: "database error: " + err.Error(),
		}
		return res
	}

	if !CheckPasswordHash(password, user.Password) {
		res := structs.LoginRes{
			Status: "fail",
			Message: "wrong password",
			Code: http.StatusUnauthorized,
		}
		return res
	}

	res := structs.LoginRes{
		Status: "success",
		Message: "loged in successfully",
		UserName: user.UserName,
		Code: http.StatusOK,
	}
	return res
}

func RegisterUser(user models.Users) structs.RegistrationRes {
    var newUser models.Users = user
    var existingUser models.Users

    // ユーザーがすでに登録されている確認
    if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
        if err != gorm.ErrRecordNotFound {
            res := structs.RegistrationRes{
                Status: "fail",
                Message: fmt.Sprintf("database error: %v", err),
				Code: http.StatusInternalServerError,
            }
            return res
        }
    } else {
        res := structs.RegistrationRes{
            Status: "fail",
            Message: "user already registered",
			Code: http.StatusConflict,
        }
        return res
    }

    if err := db.Create(&newUser).Error; err != nil {
        res := structs.RegistrationRes{
            Status: "fail",
            Message: "registration failed",
			Code: http.StatusInternalServerError,
        }
        return res
    }

    res := structs.RegistrationRes{
        Status: "success",
        Message: "registered user successfully",
		UserName: newUser.UserName,
		Code: http.StatusOK,
    }

    return res
}

func DeleteUser(email, password string) structs.DeleteRes {
	var user models.Users

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			res := structs.DeleteRes{
				Status: "fail",
				Message: "user not found",
				Code: http.StatusNotFound,
			}
			return res
		} else {
			res := structs.DeleteRes{
				Status: "fail",
				Message: "database error: " + err.Error(),
				Code: http.StatusInternalServerError,
			}
			return res
		}
	}

	if !CheckPasswordHash(password, user.Password) {
		res := structs.DeleteRes{
			Status: "fail",
			Message: "wrong password",
			Code: http.StatusUnauthorized,
		}
		return res
	}

	if err := db.Delete(&user).Error; err != nil {
		res := structs.DeleteRes{
			Status:  "fail",
			Message: fmt.Sprintf("could not delete user: %v", err),
			Code:    http.StatusInternalServerError,
		}
		return res
	}

	res := structs.DeleteRes{
		Status: "success",
		Message: fmt.Sprintf("%s was deleted successfully", user.UserName),
		Code: http.StatusOK,
	}

	return res
}