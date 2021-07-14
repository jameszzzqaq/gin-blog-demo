package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yu1er/gin-blog/pkg/e"
	"github.com/yu1er/gin-blog/pkg/utils"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = e.SUCCESS

		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			_, err := utils.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
