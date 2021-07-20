package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/junminhong/ohi-chat/config"
	"github.com/junminhong/ohi-chat/utils"
)

type MemberInfo struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func memberRegister(c *gin.Context) {
	memberInfo := MemberInfo{}
	c.BindJSON(&memberInfo)
	fmt.Println(memberInfo.Account)
	if memberInfo.Account == "" {
		fmt.Println("空")
	}
	fmt.Println(memberInfo.Account)
	c.JSON(http.StatusOK, gin.H{
		"message": "幫你註冊會員拉",
	})
}

type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Age  uint8
}

func main() {
	utils.ShowSystemMsg("oHi-chat server 正在努力啟動中！")
	utils.ShowSystemMsg("正在進行資料庫設定！")
	db := config.SetupDB()
	db.AutoMigrate(&User{})
	user := User{Name: "Jinzhu", Age: 18}
	result := db.Create(&user)
	fmt.Println(result)
}
