package models

type Order struct {
	// gorm.Model
	ID          int    `json:"id" gorm:"column:id"`
	ProductName string `json:"product_name" gorm:"column:product_name"`
	Status      string `json:"status" gorm:"column:status"`
}
