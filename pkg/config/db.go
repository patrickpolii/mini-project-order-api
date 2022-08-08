package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbURL := "postgres://postgres:c798pafo@localhost:5432/test_order"
	database, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	// database.AutoMigrate(&models.Order{})

	DB = database
}
