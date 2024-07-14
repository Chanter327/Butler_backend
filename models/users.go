package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
    gorm.Model
	UserId string
    UserName string
    Email string
	Password string
	RegisteredAt time.Time
}
