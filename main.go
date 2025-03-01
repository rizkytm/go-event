package main

import (
	"example.com/go-event/db"
	"example.com/go-event/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
