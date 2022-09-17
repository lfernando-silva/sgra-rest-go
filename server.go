package main

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lfernando-silva/sgra-rest-go/routes/index"
	"github.com/lfernando-silva/sgra-rest-go/database"
)

func getVar(envVarStr, def string) string{
	value, exists := os.LookupEnv(envVarStr)
	if !exists {
		value = def
	}
	return value
}

func loadEnv() {
	appEnv := getVar("APP_ENV", "development")
	err := godotenv.Load(".env."+appEnv, ".env.development")
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
}

func loadDatabase(){
	database.Init()
	database.AutoMigrate()
}

func main() {
	loadEnv()
	loadDatabase()

	server := gin.Default()
	appPort := getVar("APP_PORT", ":5068")

	// /
	index.HealthcheckRoutes(server)
	index.SessionRoutes(server);

	server.Run(appPort)
}