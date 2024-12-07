package config

import (
	"fmt"
	"go-todo/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var DB  *gorm.DB

func ConnectDatabase(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file") //error tp tetep jalan
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// dsn := "root:root@tcp(127.0.0.1:3306)/golang_todo?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, dbHost, dbPort, dbName)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database") //error aplikasi berhenti
	}

	DB.AutoMigrate(models.Todo{})
}