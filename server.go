package main

import (
	"log"
	"os"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

func main() {
	loadEnv()

	server := gin.Default()
	appPort := getVar("APP_PORT", ":5068")

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK",
			"timestamp": time.Now().Unix(),
		})
	})

	server.Run(appPort)
}