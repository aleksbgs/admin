package database

import (
	"../models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=go_admin port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}
	migration := database.AutoMigrate(&models.User{}, models.Role{}, models.Permission{})

	DB = database

	if migration != nil {
		panic("error with migration ")
	}
}
