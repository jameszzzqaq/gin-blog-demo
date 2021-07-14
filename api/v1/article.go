package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yu1er/gin-blog/config"
	"github.com/yu1er/gin-blog/model"
	"github.com/yu1er/gin-blog/pkg/e"
	"github.com/yu1er/gin-blog/pkg/utils"
)

func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})

	data["list"] = model.GetArticlesPage(utils.GetPage(c), config.PageSize)
	data["total"] = model.GetArticlesCount()

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

	if !model.CheckArticlExistById(id) {
		code = e.ERROR_ARTICLE_NOT_EXIST
	} else {
		code = e.SUCCESS
		article = model.GetArticleById(id)
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

	model.AddArticle(article)

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

	model.UpdateArticle(id, article)

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

	if !model.CheckArticlExistById(id) {
		code = e.ERROR_ARTICLE_NOT_EXIST
	} else {
		code = e.SUCCESS
		model.DeleteArticle(id)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})

}
