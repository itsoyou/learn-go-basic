package main

import (
	"example.com/event-booking/db"
	"example.com/event-booking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	// Setup an Engine(http server) + logger + recovery middleware
	server := gin.Default() // returns *gin.Engine

	routes.RegisterRoutes(server)

	server.Run(":8080")
}