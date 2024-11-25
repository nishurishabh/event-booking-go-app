package main

import (
	"example.com/event-booking-backend-app/db"
	"example.com/event-booking-backend-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
