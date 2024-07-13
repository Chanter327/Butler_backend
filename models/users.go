package database

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
    gorm.Model
	UserId string
    UserName string
    Email string
	RegisteredAt time.Time
}
