package main

import (
	

	"github.com/abdallahelassal/go-jwt-authentication-.git/helpers"
	"github.com/abdallahelassal/go-jwt-authentication-.git/database"
	"github.com/gin-gonic/gin"
)

func main(){
	helpers.LoadConfig(".env")
	database.ConnectDatabase()
	r := gin.Default()
	
	r.Run(":"+ helpers.AppConfig.PORT)
}