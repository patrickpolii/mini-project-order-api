package main

import (
	"order-api/pkg/config"
	"order-api/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	routes := gin.Default()

	// Connect to database
	config.ConnectDatabase()

	// Routes
	routes.GET("/orders", controllers.FindOrders)
	routes.GET("/orders/:id", controllers.FindOrder)
	routes.POST("/orders", controllers.CreateOrder)
	routes.PATCH("/orders/:id", controllers.UpdateOrder)
	routes.DELETE("/orders/:id", controllers.DeleteOrder)

	// Run the server

	routes.Run()

}
