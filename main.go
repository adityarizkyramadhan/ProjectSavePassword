package main

import (
	"ProjectSavePassword/config"
	"ProjectSavePassword/handler"
	"ProjectSavePassword/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	_, err := config.InitDB()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://save-password-frontend-26u2pyf12-adityarizkyramadhan.vercel.app", "https://save-password-frontend.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"Content": "Hello World"})
	})
	r.POST("/login", handler.LoginUser)
	r.POST("/register", handler.RegisterUser)
	r.POST("/adddata", middleware.MiddlewareJWT(), handler.AddData)
	r.GET("/seedata", middleware.MiddlewareJWT(), handler.SearchData)
	// port := fmt.Sprintf(":%d", os.Getenv("PORT"))
	r.Run()

}
