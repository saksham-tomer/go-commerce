package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func EnvMongoURI()string{
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file for mongoDB")
	}
	return os.Getenv("MONGO_URI")
}