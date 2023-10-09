package configs

import (
	"github.com/joho/godotenv"
	"log"
)

type Configuration struct {
	TorConf     TorConfig
	DbConf      DBConfig
	LoggerConf  LoggerConfig
	CrawlerConf CollyConfig
}

var cfg Configuration

func GetConfig() *Configuration {
	return &cfg
}

func SetupEnvironment() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	SetupTor()
	SetupDatabase()
	SetupLogger()
	LoadCrawlerConfig()

}
