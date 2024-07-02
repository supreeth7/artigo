package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/supreeth7/artigo/handlers"
)

func Register(router *gin.Engine) {
	router.GET("/articles/:id", handlers.GetArticleByID)
	router.GET("/articles", handlers.GetArticles)
	router.POST("/articles", handlers.CreateArticle)
}
