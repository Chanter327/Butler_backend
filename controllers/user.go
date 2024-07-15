package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	models "github.com/Chanter327/Butler_backend/models"
	services "github.com/Chanter327/Butler_backend/services"
	structs "github.com/Chanter327/Butler_backend/structs"
)

func Login(c *gin.Context) {
	var loginReq structs.LoginReq

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, structs.LoginRes{
			Status: "fail",
			Message: "invalid request: " + err.Error(),
		})
		return
	}

	result := services.Authentication(loginReq.Email, loginReq.Password)
	c.JSON(result.Code, result)
}

func NewRegistration(c *gin.Context) {
	var registrationReq structs.RegistrationReq

	if err := c.ShouldBindJSON(&registrationReq); err != nil {
		c.JSON(http.StatusBadRequest, structs.RegistrationRes{
			Status: "fail",
			Message: "invalid request: " + err.Error(),
		})
		return
	}

	userId := uuid.New().String()
	registeredAt := time.Now()
	password, err := services.HashPassword(registrationReq.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, structs.RegistrationRes{
			Status: "fail",
			Message: "invalid password: " + err.Error(),
		})
		return
	}

	newUser := models.Users{
		UserId: userId,
		UserName: registrationReq.UserName,
		Email: registrationReq.Email,
		Password: password,
		RegisteredAt: registeredAt,
	}

	result := services.RegisterUser(newUser)
	c.JSON(result.Code, result)
}

func DeleteUser(c *gin.Context) {
	var deleteReq structs.DeleteReq

	if err := c.ShouldBindJSON(&deleteReq); err != nil {
		c.JSON(http.StatusBadRequest, structs.DeleteRes{
			Status: "fail",
			Message: "invalid request: " + err.Error(),
		})
		return
	}

	result := services.DeleteUser(deleteReq.Email, deleteReq.Password)

	c.JSON(result.Code, result)
}