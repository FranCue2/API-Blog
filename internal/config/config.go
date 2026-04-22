package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	IsProduction bool

	DataBaseURI string

	AdminEmail string
	AdminPassword string
	JWTKey string

	FrontHost string
	Host string
	Port string
}

func Load() (*AppConfig, error) {

	isProduction := os.Getenv("APP_ENV") == "production" 
	
	if (!isProduction){
		err := loadEnv()
		if err != nil {
			return nil, err
		}
	}

	return &AppConfig{
		IsProduction: isProduction,

		DataBaseURI: os.Getenv("MONGO_URI"),
		AdminEmail: os.Getenv("ADMIN_EMAIL"),
		AdminPassword: os.Getenv("ADMIN_PASSWORD"),
		JWTKey: os.Getenv("JWT_SECRET"),

		FrontHost: os.Getenv("FRONT_END_HOST"),
		Host: os.Getenv("HOST"),
		Port: getEnvOrDefault("PORT", "8080"),
	}, nil
}

func getEnvOrDefault(key string, defaultValue string) string{
	res := os.Getenv(key)
	
	if res == "" {
		return defaultValue
	}

	return res
}

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return ErrCouldNotReadEnv
	}
	return nil
}