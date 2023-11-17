package config

import (
	"swisscast-go/utils"

	"github.com/gin-gonic/gin"
)

var config *Config

type Config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	gin.Default()
}

func GetDB() DBConfig {
	dbConfig := DBConfig{
		Host:    utils.GoDotEnvVariable("DB_HOST"),
		Port:     utils.GoDotEnvVariable("DB_PORT"),
		User:     utils.GoDotEnvVariable("DB_USER"),
		Pass:     utils.GoDotEnvVariable("DB_PASS"),
		Database: utils.GoDotEnvVariable("DB_DATABASE"),
	}

	return dbConfig
}
