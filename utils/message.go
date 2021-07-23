package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 200
	ERROR   = 404
)

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",
}

type CTmp struct {
	C *gin.Context
}

func ShowSystemMsg(msg string) {
	log.Println("系統訊息：" + msg)
}

func ShowErrSystemMsg(msg string) {
	log.Fatalln("系統訊息：" + msg)
}

func (cTmp CTmp) ResponseMsg(statusCode int, resultCode int, data interface{}) {
	cTmp.C.JSON(statusCode, gin.H{
		"result":  resultCode,
		"message": MsgFlags[resultCode],
		"data":    data,
	})
}
