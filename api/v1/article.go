package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yu1er/gin-blog/config"
	"github.com/yu1er/gin-blog/model"
	"github.com/yu1er/gin-blog/pkg/e"
	"github.com/yu1er/gin-blog/pkg/utils"
	"github.com/yu1er/gin-blog/service"
)

func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})

	data["list"] = service.GetArticlesPage(utils.GetPage(c), config.PageSize)
	data["total"] = service.GetArticlesCount()

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": data,
	})
}

func GetArticleById(c *gin.Context) {
	var code int
	var article model.Article
	id, _ := strconv.Atoi(c.Param("id"))

	if !service.CheckArticlExistById(id) {
		code = e.ERROR_ARTICLE_NOT_EXIST
	} else {
		code = e.SUCCESS
		article = service.GetArticleById(id)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": article,
	})

}

func AddArticle(c *gin.Context) {
	tagId, _ := strconv.Atoi(c.Query("tag_id"))
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")

	article := model.Article{
		TagID:      tagId,
		Title:      title,
		Desc:       desc,
		Content:    content,
		CreatedBy:  createdBy,
		ModifiedBy: createdBy,
	}

	service.AddArticle(article)

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		// "data": article,
	})
}

func UpdateArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tagId, _ := strconv.Atoi(c.Query("tag_id"))
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	article := model.Article{
		TagID:      tagId,
		Title:      title,
		Desc:       desc,
		Content:    content,
		ModifiedBy: modifiedBy,
	}

	service.UpdateArticle(id, article)

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		// "data": article,
	})
}

func DeleteArticle(c *gin.Context) {
	var code int
	id, _ := strconv.Atoi(c.Param("id"))

	if !service.CheckArticlExistById(id) {
		code = e.ERROR_ARTICLE_NOT_EXIST
	} else {
		code = e.SUCCESS
		service.DeleteArticle(id)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})

}
