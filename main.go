package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/supreeth7/artigo/database"
	"github.com/supreeth7/artigo/routes"
)

func main() {
	server := gin.Default()
	routes.Register(server)

	// Initialize database connection
	if err := database.Init(); err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
