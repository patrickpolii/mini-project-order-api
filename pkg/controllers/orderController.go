package controllers

import (
	"net/http"

	"order-api/pkg/config"
	"order-api/pkg/helper"
	"order-api/pkg/models"

	"github.com/gin-gonic/gin"
)

type NewOrder struct {
	ProductName string `json:"product_name" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type OrderUpdate struct {
	ProductName string `json:"product_name"`
	Status      string `json:"status"`
}

// GET /orders
// Find all orders
func FindOrders(c *gin.Context) {
	p := helper.GeneratePaginationRequest(c)
	//tr = total rows
	tr := 0
	offset := p.Page * p.Limit

	var orders []models.Order
	errFind := config.DB.Limit(p.Limit).Offset(offset).Find(&orders).Error
	if errFind != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": errFind})
	}
	p.Rows = orders

	totalRows := int64(tr)
	errCount := config.DB.Model(orders).Count(&totalRows).Error
	if errCount != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": errCount})
	}

	p.TotalRows = totalRows

	c.JSON(http.StatusOK, gin.H{"data": p})
}

// GET /orders/:id
// Find a order
func FindOrder(c *gin.Context) {
	// Get model if exist
	var order models.Order
	if err := config.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// POST /orders
// Create new order
func CreateOrder(c *gin.Context) {
	// Validate input
	var input NewOrder
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create order
	order := models.Order{ProductName: input.ProductName, Status: input.Status}
	config.DB.Create(&order)

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// PUT /orders/:id
// Update a order
func UpdateOrder(c *gin.Context) {
	// Get model if exist
	var order models.Order
	if err := config.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input OrderUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&order).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// DELETE /orders/:id
// Delete a order
func DeleteOrder(c *gin.Context) {
	// Get model if exist
	var order models.Order
	if err := config.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Delete(&order)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
