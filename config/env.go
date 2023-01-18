package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvOpenAI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	return os.Getenv("OPENAI_API_KEY")
}
