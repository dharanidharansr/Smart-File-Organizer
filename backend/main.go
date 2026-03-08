package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to MongoDB
	ConnectDB()

	// Set up Gin router
	r := gin.Default()

	// CORS – allow the Astro dev server (port 4321) and any localhost origin
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4321", "http://127.0.0.1:4321"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: false,
	}))

	// Register REST routes
	RegisterRoutes(r)

	// Start on port 8081
	r.Run(":8081")
}
