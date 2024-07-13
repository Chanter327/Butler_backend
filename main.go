package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"./models/users"
)

func main() {
	user := showUser()
	userName := user.UserName

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "butler application",
		"first-user": userName,
	  })
	})
	r.Run()
}

func showUser() users.Users {
	dsn := "host=localhost user=coffee password=pass dbname=butler_db port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

	db.AutoMigrate(&users.Users{})

	var user users.Users
	db.First(&user, 1)

	return user
}