package member

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var wsList = []*websocket.Conn{}

func WsPing(ctx *gin.Context) {
	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	wsList = append(wsList, ws)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		// 讀取ws Socket傳來的訊息
		mt, message, err := ws.ReadMessage()
		fmt.Println(len(wsList))
		if err != nil {
			break
		}
		// 如果是ping
		/*if string(message) == "ping" {
			// 就回pong
			message = []byte("pong")
		} else {
			// 如果是其他, 就回文字訊息類型, 內容就是回聲 (鸚鵡XD)
			ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintln("got it : "+string(message))))
		}*/
		// 寫入Websocket
		for _, wsa := range wsList {
			pHandelr := wsa.PingHandler()
			fmt.Println(pHandelr)
			err = wsa.WriteMessage(mt, message)
			if err != nil {
				break
			}
		}
	}
}
