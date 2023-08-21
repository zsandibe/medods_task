package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURL       string
	NameDB         string
	NameCollection string
	AccessKey      string
	RefreshKey     string
}

func Configuration() Config {
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		fmt.Printf("Error loading .env file : %v\n", err)
	}
	return Config{
		MongoURL:       os.Getenv("MONGODB_URL"),
		NameDB:         os.Getenv("NAME_DB"),
		NameCollection: os.Getenv("NAME_COLLECTION"),
		AccessKey:      os.Getenv("ACCESS_KEY"),
	}
}
