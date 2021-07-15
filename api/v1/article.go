package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yu1er/gin-blog/model/request"
	"github.com/yu1er/gin-blog/model/response"
	"github.com/yu1er/gin-blog/pkg/e"
	"github.com/yu1er/gin-blog/service"
)

func GetArticles(c *gin.Context) {
	var req request.ArticleListGet

	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.Code(e.INVALID_PARAMS, c)
		return
	}

	list, total, err := service.GetArticlesPage(req)
	if err != nil {
		response.Code(e.ERROR_ARTICLE_NOT_EXIST, c)
	} else {
		response.OKWithData(response.PageResult{
			List:     list,
			Total:    total,
			PageNum:  req.PageNum,
			PageSize: req.PageSize,
		}, c)
	}
}

func GetArticleById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Code(e.INVALID_PARAMS, c)
		return
	}

	art, err := service.GetArticleById(id)
	if err != nil {
		response.Code(e.ERROR_ARTICLE_NOT_EXIST, c)
	} else {
		response.OKWithData(art, c)
	}
}

func AddArticle(c *gin.Context) {
	var req request.AriticleAdd

	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.Code(e.INVALID_PARAMS, c)
		return
	}

	err = service.AddArticle(&req)
	if err != nil {
		response.Code(e.ERROR_ARTICLE_NOT_EXIST, c)
	} else {
		response.OK(c)
	}
}

func UpdateArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Code(e.INVALID_PARAMS, c)
		return
	}

	var req request.ArticleUpdate
	err = c.ShouldBindJSON(&req)
	if err != nil {
		response.Code(e.ERROR_ARTICLE_NOT_EXIST, c)
		return
	}

	err = service.UpdateArticle(id, &req)
	if err != nil {
		response.Code(e.ERROR_ARTICLE_NOT_EXIST, c)
	} else {
		response.OK(c)
	}
}

func DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Code(e.INVALID_PARAMS, c)
		return
	}

	err = service.DeleteArticle(id)
	if err != nil {
		response.Code(e.ERROR_ARTICLE_NOT_EXIST, c)
	} else {
		response.OK(c)
	}
}
