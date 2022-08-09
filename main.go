package main

import (
	"order-api/pkg/config"
	"order-api/pkg/controllers"
	"order-api/pkg/helper"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//load env
	godotenv.Load()

	routes := gin.Default()

	// Connect to database
	config.ConnectDatabase()

	// Routes Orders
	routes.GET("/orders", helper.JwtAuthMiddleware(), controllers.FindOrders)
	routes.GET("/orders/:id", helper.JwtAuthMiddleware(), controllers.FindOrder)
	routes.POST("/orders", helper.JwtAuthMiddleware(), controllers.CreateOrder)
	routes.PATCH("/orders/:id", helper.JwtAuthMiddleware(), controllers.UpdateOrder)
	routes.DELETE("/orders/:id", helper.JwtAuthMiddleware(), controllers.DeleteOrder)

	//Routes Auth
	routes.POST("/register", controllers.Register)
	routes.POST("/login", controllers.Login)

	// Run the server

	routes.Run()

}
