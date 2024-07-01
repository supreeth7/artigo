package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/articles", getArticles)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func getArticles(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "article-A",
	})
}
