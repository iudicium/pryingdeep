package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
)

const projectDirName = "pryingdeep" // change to relevant project name

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

// Load the setup dynamically, so we can use it for tests later on too
func LoadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
func SetupEnvironment() {
	LoadEnv()
	SetupTor()
	SetupDatabase()
	SetupLogger()
	LoadCrawlerConfig()

}
