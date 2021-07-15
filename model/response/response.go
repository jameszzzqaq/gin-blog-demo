package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yu1er/gin-blog/pkg/e"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func result(code int, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{code, e.GetMsg(code), data})
}

func Code(code int, c *gin.Context) {
	result(code, nil, c)
}

func CodeWithData(code int, data interface{}, c *gin.Context) {
	result(code, data, c)
}

func OK(c *gin.Context) {
	result(e.SUCCESS, nil, c)
}

func OKWithData(data interface{}, c *gin.Context) {
	result(e.SUCCESS, data, c)
}
