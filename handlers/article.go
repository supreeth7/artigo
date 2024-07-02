package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/supreeth7/artigo/database"
	"github.com/supreeth7/artigo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateArticle(ctx *gin.Context) {
	var article models.Article

	article.ID = primitive.NewObjectID()
	article.DateTime = time.Now()

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

func UpdateArticle(ctx *gin.Context) {
	id := ctx.Param("id")

	var article models.Article

	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := article.Update(id, &database.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result.MatchedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "No record found to modify",
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func DeleteArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	var a models.Article

	res, err := a.Delete(id, &database.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if res.DeletedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "No record found to delete",
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
