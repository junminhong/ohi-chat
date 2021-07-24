package main

import (
	"crypto/rand"
	"crypto/rsa"

	routes "github.com/junminhong/ohi-chat/config/routes"
	"github.com/junminhong/ohi-chat/utils"
)

var secretKey, _ = rsa.GenerateKey(rand.Reader, 1024)

func main() {
	utils.ShowSystemMsg("oHi-chat server 正在努力啟動中！")
	routes.InitRouter()

	/*newCompany := models.Company{}
	postgresDB.Table("companies").Where("name = ?", "測試公司").Scan(&newCompany)
	postgresDB.Create(&models.Member{Name: "b", Age: 18, CompanyID: newCompany.ID, Company: newCompany})*/

}
