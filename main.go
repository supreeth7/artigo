package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/supreeth7/artigo/database"
	"github.com/supreeth7/artigo/middleware"
	"github.com/supreeth7/artigo/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file is missing")
	}

	server := gin.Default()
	server.ForwardedByClientIP = true
	server.SetTrustedProxies([]string{"127.0.0.1"})

	protected := server.Group("/")
	protected.Use(middleware.EnsureValidToken())

	routes.RegisterProtectedRoutes(protected)
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
