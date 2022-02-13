package main

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnvOrFile(key string) (envValue string) {
	envValue = os.Getenv(key)
	if envValue == "" {
		currentEnv, err := godotenv.Read("../.env")
		if err != nil {
			return ""
		}
		return currentEnv[key]
	}
	return envValue
}
