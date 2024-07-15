// routes/router.go
package routes

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	controllers "github.com/Chanter327/Butler_backend/controllers"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // フロントエンドのURLを指定
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/api/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "butler application",
		})
	})

	r.POST("/api/login", func(c *gin.Context) {
		controllers.Login(c)
	})
	r.POST("/api/user/registration", func(c *gin.Context) {
		controllers.NewRegistration(c)
	})
	r.POST("/api/user/delete", func(c *gin.Context) {
		controllers.DeleteUser(c)
	})

	return r
}
