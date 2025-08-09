package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func Get() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file: ", err.Error())
	}

	jwtExp, err := strconv.Atoi(os.Getenv("JWT_EXP"))
	if err != nil {
		jwtExp = 10
	}

	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: Database{
			Host: os.Getenv("DATABASE_HOST"),
			Port: os.Getenv("DATABASE_PORT"),
			User: os.Getenv("DATABASE_USER"),
			Pass: os.Getenv("DATABASE_PASS"),
			Name: os.Getenv("DATABASE_NAME"),
			Tz:   os.Getenv("DATABASE_TZ"),
		},
		Jwt: Jwt{
			Key: os.Getenv("JWT_KEY"),
			Exp: jwtExp,
		},
	}
}
