package main

import (
	"log"
	"os"
	"time"

	"github.com/MechamJonathan/palantir-cli/internal/theoneapi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY not set in .env file")
	}

	theoneClient := theoneapi.NewClient(5*time.Second, time.Minute*5, apiKey)
	cfg := &config{
		theoneapiClient: theoneClient,
	}

	startRepl(cfg)
}
