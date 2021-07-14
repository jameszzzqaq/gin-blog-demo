package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/yu1er/gin-blog/model"
	"github.com/yu1er/gin-blog/pkg/e"
	"github.com/yu1er/gin-blog/pkg/utils"
	"github.com/yu1er/gin-blog/service"
)

func GetAuth(c *gin.Context) {
	data := make(map[string]interface{})
	v := validator.New()
	username := c.Query("username")
	password := c.Query("password")

	auth := model.Auth{Username: username, Password: password}

	var code int
	err := v.Struct(auth)
	if err != nil {
		code = e.INVALID_PARAMS
		goto response
	}

	if exist := service.CheckAuth(username, password); exist {
		token, err := utils.GenerateToken(username)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token
			code = e.SUCCESS
		}
	}

response:
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
