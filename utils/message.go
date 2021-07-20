package utils

import "log"

func ShowSystemMsg(msg string) {
	log.Println("系統訊息：" + msg)
}

func ShowErrSystemMsg(msg string) {
	log.Fatalln("系統訊息：" + msg)
}
