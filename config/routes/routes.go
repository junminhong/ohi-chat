package routes

import (
	"github.com/gin-gonic/gin"
	chatroom "github.com/junminhong/ohi-chat/app/controllers/chat-room"
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
	chatRoomRouter := v1.Group("/chat-room")
	{
		//chatRoom.GET("/", chat.WsPing)
		chatRoomRouter.POST("/create", chatroom.Create)
		chatRoomRouter.PUT("/edit", chatroom.Edit)
		chatRoomRouter.POST("/join", chatroom.Join)
		chatRoomRouter.POST("/leave", chatroom.Leave)
		chatRoomRouter.GET("/list", chatroom.List)
		chatRoomRouter.DELETE("/delete", chatroom.Delete)
	}

	utils.ShowSystemMsg("oHi-chat server 正在努力運行中！")
	router.Run(":8070")
}
