package main

import (
	"github.com/gin-gonic/gin"
	"main.go/handlers"
	"main.go/initializer"
)

func init() {
	initializer.Initailize()
}
func main() {

	r := gin.Default()

	r.POST("/signup", handlers.SignUp)
	r.GET("/login", handlers.Login)
	r.DELETE("/logout", handlers.Logout)

	r.Run(":8080")
}
