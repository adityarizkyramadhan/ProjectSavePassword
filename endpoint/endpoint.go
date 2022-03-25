package endpoint

import (
	"ProjectSavePassword/handler"
	"ProjectSavePassword/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartEndpoint() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/login", handler.LoginUser)
	r.POST("/register", handler.RegisterUser)
	r.POST("/adddata", middleware.MiddlewareJWT(), handler.AddData)
	r.GET("/seedata", middleware.MiddlewareJWT(), handler.SearchData)
	//port := fmt.Sprintf(":%d", os.Getenv("PORT"))
	r.Run(":48840")
}
