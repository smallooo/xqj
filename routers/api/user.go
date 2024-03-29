package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"xqj/models"
	"xqj/pkg/e"
	"xqj/pkg/logging"
	"xqj/pkg/util"
	"xqj/utils"
)

type authAccount struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
	PhoneNo  string `valid:"Required; MaxSize(50)"`
}

// login
// @Summary Login user
// @Produce  json
// @Param username  query string true "username"
// @Param password  query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func CreateUser(c *gin.Context) {
	phoneno := c.Query("phoneno")
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := authAccount{Username: username, Password: password, PhoneNo: phoneno}
	ok, _ := valid.Valid(&a)
	if ok {
		ok = utils.VerifyMobileFormat(phoneno)
	}

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = e.SUCCESS
			}

			//models.LoginRecord()
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
