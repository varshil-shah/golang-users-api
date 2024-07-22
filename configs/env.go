package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(variableName string) (string, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error while fetching environmental variables.")
	}

	result := os.Getenv(variableName)
	if result == "" {
		return "", fmt.Errorf("environment variable %s not found", variableName)
	}

	return result, nil
}
