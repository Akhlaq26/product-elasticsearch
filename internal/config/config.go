package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
	ServerPort  string
	EsUser      string
	EsHost      string
	EsPort      string
	EsPass      string
}

func Init() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf(".env not found, load env from operating system only")
	}

	// init, check data and set default value
	config := &Config{}

	if environment := os.Getenv("ENVIRONMENT"); len(environment) > 0 {
		config.Environment = environment
	}

	if serverPort := os.Getenv("SERVER_PORT"); len(serverPort) > 0 {
		config.ServerPort = serverPort
	}

	if esHost := os.Getenv("ES_HOST"); len(esHost) > 0 {
		config.EsHost = esHost
	}

	if esPort := os.Getenv("ES_PORT"); len(esPort) > 0 {
		config.EsPort = esPort
	}

	if esUser := os.Getenv("ES_USER"); len(esUser) > 0 {
		config.EsUser = esUser
	}

	if esPass := os.Getenv("ES_PASS"); len(esPass) > 0 {
		config.EsPass = esPass
	}

	return config
}
