package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	routes "github.com/newlinedeveloper/go-boilerplate/routes"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": "Access granted for api-2"})
	})


	router.Run(":8000")
}
