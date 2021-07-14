package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yu1er/gin-blog/config"
)

func GetPage(c *gin.Context) int {
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}

	return (page - 1) * config.PageSize
}
