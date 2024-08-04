package main

import (
	"github.com/Amit9887/go-jwt-authentication/controllers"
	"github.com/Amit9887/go-jwt-authentication/initializers"
	"github.com/Amit9887/go-jwt-authentication/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
