package member

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/junminhong/ohi-chat/app/models"
	database "github.com/junminhong/ohi-chat/config/database"
	"github.com/junminhong/ohi-chat/utils"
)

type memberStruct struct {
	Account  string
	Password string
	Token    string
}

func createMember(account string, password string) bool {
	postgresDB := database.GetDB()
	result := postgresDB.Create(&models.Member{Account: account, Password: password})
	return result.Error == nil
}

func Signup(c *gin.Context) {
	memberStruct := &memberStruct{}
	resp := utils.CTmp{C: c}
	if c.BindJSON(memberStruct) != nil {
		resp.ResponseMsg(http.StatusOK, 404, nil)
	}
	if memberStruct.Account == "" || memberStruct.Password == "" {
		resp.ResponseMsg(http.StatusOK, 404, nil)
	} else {
		if createMember(memberStruct.Account, memberStruct.Password) {
			resp.ResponseMsg(http.StatusOK, 200, nil)
		} else {
			resp.ResponseMsg(http.StatusOK, 404, nil)
		}
	}
}

func Signin(c *gin.Context) {
	// 簽發一個access_token和refresh_token給front-end
	// 這兩個token要存回db
	// redis
	nowWorkDir, err := os.Getwd()
	if err != nil {
		utils.ShowErrSystemMsg(err.Error())
	}
	SECRETKEY, err := ioutil.ReadFile(nowWorkDir + "/key.pem")
	if err != nil {
		utils.ShowErrSystemMsg(err.Error())
	}
	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}
	now := time.Now()
	jwtID := strconv.FormatInt(now.Unix(), 10)
	fmt.Println(now.Add(20 * time.Second).Unix())
	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			Audience:  "test",
			ExpiresAt: now.Add(60 * time.Second).Unix(),
			Id:        jwtID,
			IssuedAt:  now.Unix(),
			Issuer:    "ginJWT",
			//NotBefore: now.Add(20 * time.Second).Unix(),
			Subject: "testa",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(SECRETKEY))
	if err != nil {
		utils.ShowErrSystemMsg(err.Error())
	}
	access_token, err := token.SignedString(privateKey)
	if err != nil {
		utils.ShowErrSystemMsg(err.Error())
	}
	fmt.Println(access_token)
}

func Test(c *gin.Context) {

}
