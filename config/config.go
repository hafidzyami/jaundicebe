package config

import (
	"log"
	"github.com/spf13/viper"
	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	viper.AutomaticEnv()
}

func GetEnv(key string) string {
	if value := viper.GetString(key); value != "" {
		return value
	}
	return ""
}