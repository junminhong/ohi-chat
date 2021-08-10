package chat

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Member struct {
	Websocket *websocket.Conn
	NickName  string
	Activity  bool
}

var memberList = []Member{}

func WsPing(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	newMember := Member{Websocket: ws, NickName: "", Activity: true}
	memberList = append(memberList, newMember)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		// 讀取ws Socket傳來的訊息
		mt, msgTmp, err := ws.ReadMessage()
		if err != nil {
			return
		}
		// 寫入Websocket
		for _, member := range memberList {
			log.Println(member.Activity)
			if member.Activity {
				if member.Websocket == ws {
					continue
				}
				msg := member.NickName + string(msgTmp)
				err = member.Websocket.WriteMessage(mt, []byte(msg))
				if err != nil {
					member.Activity = false
					log.Println(err.Error())
				}
			}
		}
	}
}
