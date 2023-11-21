package routes

import (
	controller "github.com/newlinedeveloper/go-boilerplate/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(authRoutes *gin.Engine) {
	authRoutes.POST("users/signup", controller.SignUp())

}