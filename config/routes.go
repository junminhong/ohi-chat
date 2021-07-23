package config

import (
	"github.com/gin-gonic/gin"
	"github.com/junminhong/ohi-chat/app/controllers"
)

func InitRouter() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/signup", controllers.Signup)
		//v1.POST("/signin", submitEndpoint)
	}

	router.Run(":8080")
}
