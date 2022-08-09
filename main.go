package main

import (
	"order-api/pkg/config"
	"order-api/pkg/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	routes := gin.Default()

	// Connect to database
	config.ConnectDatabase()

	// Routes Orders
	routes.GET("/orders", controllers.FindOrders)
	routes.GET("/orders/:id", controllers.FindOrder)
	routes.POST("/orders", controllers.CreateOrder)
	routes.PATCH("/orders/:id", controllers.UpdateOrder)
	routes.DELETE("/orders/:id", controllers.DeleteOrder)

	//Routes Auth
	routes.POST("/register", controllers.Register)
	routes.POST("/login", controllers.Login)

	// Run the server

	routes.Run()

}
