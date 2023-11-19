package config

import (
	"swisscast-go/utils/AppUtils"

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
		Host:    AppUtils.GoDotEnvVariable("DB_HOST"),
		Port:     AppUtils.GoDotEnvVariable("DB_PORT"),
		User:     AppUtils.GoDotEnvVariable("DB_USER"),
		Pass:     AppUtils.GoDotEnvVariable("DB_PASS"),
		Database: AppUtils.GoDotEnvVariable("DB_DATABASE"),
	}

	return dbConfig
}
