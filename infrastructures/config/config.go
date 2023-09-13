package config

import (
	"os"

	"github.com/joho/godotenv"
)

func New() LoadConfig {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	return LoadConfig{
		App: App{
			Mode:       os.Getenv("APP_MODE"),
			Name:       os.Getenv("APP_NAME"),
			Port:       os.Getenv("APP_PORT"),
			Url:        os.Getenv("APP_URL"),
			Secret_key: os.Getenv("SECRET_KEY"),
		},
		Database: Database{
			Host:     os.Getenv("DB_HOST"),
			Name:     os.Getenv("DB_NAME"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Port:     os.Getenv("DB_PORT"),
		},
	}
}
