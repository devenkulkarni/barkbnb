package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App struct {
		Name    string
		Env     string
		Version string
	}

	Server struct {
		Port string
	}

	Database struct {
		Host     string
		Port     string
		Name     string
		User     string
		Password string
		SSLMode  string
	}
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}
	cfg := &Config{}
	cfg.App.Name = os.Getenv("APP_NAME")
	cfg.App.Env = os.Getenv("APP_ENV")
	cfg.App.Version = os.Getenv("APP_VERSION")

	cfg.Server.Port = os.Getenv("SERVER_PORT")

	cfg.Database.Host = os.Getenv("POSTGRES_HOST")
	cfg.Database.Port = os.Getenv("POSTGRES_PORT")
	cfg.Database.Name = os.Getenv("POSTGRES_DB_NAME")
	cfg.Database.User = os.Getenv("POSTGRES_USER")
	cfg.Database.Password = os.Getenv("POSTGRES_PASSWORD")
	cfg.Database.SSLMode = os.Getenv("POSTGRES_SSL_MODE")

	if cfg.Server.Port == "" {
		return nil, fmt.Errorf("SERVER_PORT is required")
	}

	if cfg.Database.Host == "" || cfg.Database.Port == "" || cfg.Database.Name == "" || cfg.Database.User == "" || cfg.Database.Password == "" {
		return nil, fmt.Errorf("POSTGRES_HOST, POSTGRES_PORT, POSTGRES_DB_NAME, POSTGRES_USER, and POSTGRES_PASSWORD are required")
	}

	return cfg, nil
}
