package config

import (
    "os"
    "strconv"

    "github.com/joho/godotenv"
)

type Config struct {
    AppPort             string
    DBHost              string
    DBPort              string
    DBUser              string
    DBPassword          string
    DBName              string
    JWTSecret           string
    TokenDurationMinute int
}

func LoadConfig() *Config {
    _ = godotenv.Load()
    

    duration, _ := strconv.Atoi(os.Getenv("TOKEN_DURATION_MINUTES"))

    return &Config{
        AppPort:             os.Getenv("APP_PORT"),
        DBHost:              os.Getenv("DB_HOST"),
        DBPort:              os.Getenv("DB_PORT"),
        DBUser:              os.Getenv("DB_USER"),
        DBPassword:          os.Getenv("DB_PASSWORD"),
        DBName:              os.Getenv("DB_NAME"),
        JWTSecret:           os.Getenv("JWT_SECRET"),
        TokenDurationMinute: duration,
    }

	
}
