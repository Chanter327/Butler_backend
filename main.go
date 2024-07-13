package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	services "github.com/Chanter327/Butler_backend/services"
)

func main() {
	user, err := services.ShowUser()
	if err != nil {
		panic(err)
	}
	
	r := gin.Default()
	r.Use(gin.Logger())
	r.GET("/", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "butler application",
		"first-user": gin.H{
			"id": user.UserId,
			// "name": user.UserName,
			// "registered_at": user.RegisteredAt,
		},
	  })
	})
	r.Run()
}