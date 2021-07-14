package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yu1er/gin-blog/api"
	v1 "github.com/yu1er/gin-blog/api/v1"
	"github.com/yu1er/gin-blog/config"
	"github.com/yu1er/gin-blog/middleware"
)

func SetupRouter() *gin.Engine {
	e := gin.New()

	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	gin.SetMode(config.RunMode)

	e.GET("/auth", api.GetAuth)

	rg := e.Group("/api/v1")
	rg.Use(middleware.JWT())
	{
		// tag
		rg.GET("/tags", v1.GetTags)
		rg.GET("/tag/:id", v1.GetTagById)
		rg.POST("/tag", v1.AddTag)
		rg.DELETE("/tag/:id", v1.DeleteTag)
		rg.PUT("/tag/:id", v1.UpdateTag)

		// article
		rg.GET("/articles", v1.GetArticles)
		rg.GET("/article/:id", v1.GetArticleById)
		rg.POST("/article/", v1.AddArticle)
		rg.PUT("/article/:id", v1.UpdateArticle)
		rg.DELETE("/article/:id", v1.DeleteArticle)
	}

	return e
}
