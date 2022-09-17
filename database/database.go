package database

import (
	"os"
	"log"
)

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/lfernando-silva/sgra-rest-go/database/models"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	POSTGRES_DB := os.Getenv("POSTGRES_DB")
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	POSTGRES_PORT := os.Getenv("POSTGRES_PORT")
	POSTGRES_SSL := os.Getenv("POSTGRES_SSL")

	dsn := "host="+POSTGRES_HOST
	dsn += " user="+POSTGRES_USER
	dsn += " password="+POSTGRES_PASSWORD
	dsn += " dbname="+POSTGRES_DB
	dsn += " port="+POSTGRES_PORT

	if POSTGRES_SSL != "" {
		dsn += " sslmode="+POSTGRES_SSL
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	  }
	//db.LogMode(true)
	DB = db
	return DB
}

func AutoMigrate() {	
	// Migrations
	// Does not work
	DB.Exec("CREATE EXTENSION IF NOT EXISTS "+"pgcrypto"+";")

	DB.AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB {
	return DB
}