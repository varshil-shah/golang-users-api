package configs

import (
	"fmt"
	"os"
)

func GetEnvVariable(variableName string) (string, error) {

	result := os.Getenv(variableName)
	if result == "" {
		return "", fmt.Errorf("environment variable %s not found", variableName)
	}

	return result, nil
}
