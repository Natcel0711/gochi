package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ConnectionBuilder(keyword string) string {
	var url string
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}
	switch keyword {
	case "postgres":
		url = fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	}
	return url
}
