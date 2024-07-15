package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Chanter327/Butler_backend/models"
	services "github.com/Chanter327/Butler_backend/services"
)

type LoginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginRes struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	UserName string `json:"userName,omitempty"`
	Code int
}

func Login(c *gin.Context) {
	var loginReq LoginReq

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, LoginRes{
			Status: "fail",
			Message: "invalid request: " + err.Error(),
		})
		return
	}

	result := services.Authentication(loginReq.Email, loginReq.Password)
	c.JSON(result.Code, result)
}

type RegistrationReq struct {
	UserName string `json:"userName" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegistrationRes struct {
	Status string `json:"status"`
	Message  string `json:"message"`
	UserName string `json:"userName,omitempty"`
	Code int
}

func NewRegistration(c *gin.Context) {
	var registrationReq RegistrationReq

	if err := c.ShouldBindJSON(&registrationReq); err != nil {
		c.JSON(http.StatusBadRequest, RegistrationRes{
			Status: "fail",
			Message: "invalid request: " + err.Error(),
		})
		return
	}

	userId := uuid.New().String()
	registeredAt := time.Now()
	password, err := services.HashPassword(registrationReq.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, RegistrationRes{
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