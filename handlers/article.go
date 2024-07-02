package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supreeth7/artigo/database"
	"github.com/supreeth7/artigo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateArticle(ctx *gin.Context) {
	var article models.Article

	article.ID = primitive.NewObjectID()

	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := article.Create(&database.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func GetArticles(ctx *gin.Context) {
	articles, err := models.Get(&database.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, articles)
}

func GetArticleByID(ctx *gin.Context) {
	var article models.Article
	id := ctx.Param("id")

	err := article.GetByID(id, &database.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if article.ID == primitive.NilObjectID {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "No article found",
		})
		return
	}

	ctx.JSON(http.StatusOK, article)
}
