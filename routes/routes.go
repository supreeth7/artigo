package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/supreeth7/artigo/handlers"
)

func Register(router *gin.Engine) {
	router.GET("/articles", handlers.GetArticles)
	router.GET("/articles/:id", handlers.GetArticleByID)
}

func RegisterProtectedRoutes(router *gin.RouterGroup) {
	router.POST("/articles", handlers.CreateArticle)
	router.PUT("/articles/:id", handlers.UpdateArticle)
	router.DELETE("/articles/:id", handlers.DeleteArticle)
}
