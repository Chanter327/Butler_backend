package database

import (
	"gorm.io/gorm"
)

type Users struct {
    gorm.Model
	UserId string
    UserName string
    Email string
	RegisteredAt string
}
