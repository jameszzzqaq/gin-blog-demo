package utils

import (
	"github.com/yu1er/gin-blog/config"
)

func GetPage(page ...int) int {
	pageNum := 0
	pageSize := config.PageSize

	if len(page) > 0 {
		pageNum = page[0]
	}

	if len(page) == 2 {
		pageSize = page[1]
	}

	if pageNum == 0 {
		pageNum = 1
	}

	return (pageNum - 1) * pageSize
}
