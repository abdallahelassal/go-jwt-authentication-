package database

import (
	"github.com/abdallahelassal/go-jwt-authentication-.git/helpers"
	"github.com/abdallahelassal/go-jwt-authentication-.git/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		helpers.AppConfig.DB_HOST,
		helpers.AppConfig.DB_USER,
		helpers.AppConfig.DB_PASSWORD,
		helpers.AppConfig.DB_NAME,
		helpers.AppConfig.DB_PORT)

    fmt.Println("DSN:", dsn)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = database
	database.AutoMigrate(&models.User{})
}
