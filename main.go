package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/supreeth7/artigo/database"
	"github.com/supreeth7/artigo/handlers"
)

func main() {
	r := gin.Default()

	// Initialize database connection
	if err := database.Init(); err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	r.POST("/articles", handlers.CreateArticle)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
