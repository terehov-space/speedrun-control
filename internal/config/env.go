package config

import (
	"context"
	"github.com/joho/godotenv"
	"log"
)

func readConfigFile(c context.Context) {
	err := godotenv.Load(c.Value("config").(string))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
