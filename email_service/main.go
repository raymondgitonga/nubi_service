package main

import (
	"github.com/joho/godotenv"
	"github.com/raymondgitonga/nubi_service/email_service/app"
	"github.com/raymondgitonga/nubi_service/email_service/utils"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOST")
	dbname   := os.Getenv("DB_NAME")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	utils.InitDatabase(username,password,hostname,dbname)

	app.StartApp()
}
