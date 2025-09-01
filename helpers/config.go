package helpers

import (
	"log"
	"os"
	// "strconv"
	"github.com/joho/godotenv"
)

var AppConfig struct{
	PORT string 
	DB_HOST string
	DB_NAME string
	DB_PASSWORD string
	DB_PORT string
	SECRET_KEY string
	DB_USER string
}

func LoadConfig(filename string){
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatal("Error loading .env file:", err)
		return
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	// db_port , err := strconv.Atoi(os.Getenv("DB_PORT"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	AppConfig.PORT = port
	AppConfig.DB_HOST = os.Getenv("DB_HOST")
	AppConfig.DB_NAME = os.Getenv("DB_NAME")
	AppConfig.DB_PORT = os.Getenv("DB_PORT")
	AppConfig.DB_USER = os.Getenv("DB_USER")
	AppConfig.SECRET_KEY = os.Getenv("SECRET_KEY")
	
}