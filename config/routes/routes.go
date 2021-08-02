package config

import (
	"github.com/gin-gonic/gin"
	"github.com/junminhong/ohi-chat/app/controllers/member"
	"github.com/junminhong/ohi-chat/utils"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func InitRoutes() {
	utils.ShowSystemMsg("oHi-chat server router 正在努力啟動中！")
	router := gin.Default()
	router.Use(CORSMiddleware())
	v1 := router.Group("/api/v1")
	{
		v1.POST("/signup", member.Signup)
		v1.POST("/signin", member.Signin)
		v1.POST("/test", member.Test)
		v1.GET("/ping", member.WsPing)
	}
	utils.ShowSystemMsg("oHi-chat server 正在努力運行中！")
	router.Run(":8070")
}
