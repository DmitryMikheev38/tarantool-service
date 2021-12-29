package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port        string
	TarantoolDB TarantoolDB
}

type TarantoolDB struct {
	Host     string
	Port     string
	User     string
	Password string
}

func New() *Config {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	return &Config{
		Port: os.Getenv("APP_PORT"),
		TarantoolDB: TarantoolDB{
			Host:     os.Getenv("TARANTOOL_HOST"),
			Port:     os.Getenv("TARANTOOL_PORT"),
			User:     os.Getenv("TARANTOOL_USER"),
			Password: os.Getenv("TARANTOOL_PWD"),
		},
	}
}
