package member

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/junminhong/ohi-chat/app/models"
	database "github.com/junminhong/ohi-chat/config/database"
	"github.com/junminhong/ohi-chat/utils"
)

type SignupStruct struct {
	Account  string
	Password string
}

func createMember(account string, password string) bool {
	postgresDB := database.GetDB()
	result := postgresDB.Create(&models.Member{Account: account, Password: password})
	return result.Error == nil
}

func Signup(c *gin.Context) {
	signupStruct := &SignupStruct{}
	resp := utils.CTmp{C: c}

	if c.BindJSON(signupStruct) != nil {
		resp.ResponseMsg(http.StatusOK, 404, nil)
	}
	if signupStruct.Account == "" || signupStruct.Password == "" {
		resp.ResponseMsg(http.StatusOK, 404, nil)
	} else {
		if createMember(signupStruct.Account, signupStruct.Password) {
			resp.ResponseMsg(http.StatusOK, 200, nil)
		} else {
			resp.ResponseMsg(http.StatusOK, 404, nil)
		}
	}
}
