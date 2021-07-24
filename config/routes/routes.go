package config

import (
	"github.com/gin-gonic/gin"
	"github.com/junminhong/ohi-chat/app/controllers/member"
)

func InitRouter() {

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/signup", member.Signup)
		v1.POST("/signin", member.Signin)
		v1.POST("/test", member.Test)
	}

	router.Run(":8080")
}
