package configs

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	"github.com/fatih/color"
)

const projectDirName = "pryingdeep" // change to relevant project name

type Configuration struct {
	TorConf     TorConfig
	DbConf      DBConfig
	LoggerConf  LoggerConfig
	CrawlerConf CollyConfig
	PryingConf  PryingConfig
}

var cfg Configuration

func GetConfig() *Configuration {
	return &cfg
}

// Load the setup dynamically, so we can use it for tests later on too
func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func loadConfig(configFile string, config interface{}) {

	errorPrinter := color.New(color.FgRed).Printf

	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName(configFile)

	if err := viper.ReadInConfig(); err != nil {
		errorPrinter("Error during loading %s config: %s\n", configFile, err)
		return
	}

	if err := viper.Unmarshal(config); err != nil {
		errorPrinter("Error during binding %s config to struct: %s\n", configFile, err)
		return
	}
	color.HiMagenta(color.Green("[+]"Sucesfully loaded %s", configFile)
}

func SetupEnvironment() {
	loadEnv()
	setupTor()
	setupLogger()
	loadCrawlerConfig()
	loadPryingConfig()
	LoadDatabase()

}
