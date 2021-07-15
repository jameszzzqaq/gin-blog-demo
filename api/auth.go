package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yu1er/gin-blog/model"
	"github.com/yu1er/gin-blog/model/request"
	"github.com/yu1er/gin-blog/model/response"
	"github.com/yu1er/gin-blog/pkg/e"
	"github.com/yu1er/gin-blog/pkg/utils"
	"github.com/yu1er/gin-blog/service"
)

func GetAuth(c *gin.Context) {
	var authReq request.AuthReq
	err := c.ShouldBindJSON(&authReq)
	if err != nil {
		response.Code(e.ERROR_AUTH_REQ_PARAM, c)
		return
	}

	username, password := authReq.Username, authReq.Passwrod

	// 数据库查询是否存在该用户
	auth := &model.Auth{Username: username, Password: password}
	if exist := service.CheckAuthExist(auth); !exist {
		response.Code(e.ERROR_AUTH_VALIDATE_FAIL, c)
		return
	}

	// 为用户生成token
	token, err := utils.GenerateToken(username)
	if err != nil {
		response.Code(e.ERROR_AUTH_TOKEN, c)
		return
	}

	// 返回token
	data := response.AuthResp{Username: username, Token: token}
	response.OKWithData(data, c)
}
