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
	TorConf      TorConfig
	DbConf       DBConfig
	LoggerConf   LoggerConfig
	CrawlerConf  CollyConfig
	PryingConf   PryingConfig
	ExporterConf ExporterConfig
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

func loadConfig(configFile string, config interface{}) {

	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName(configFile)

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error during loading %s config: %s\n", configFile, err)
		return
	}

	if err := viper.Unmarshal(config); err != nil {
		log.Printf("Error during binding %s config to struct: %s\n", configFile, err)
		return
	}
	color.HiMagenta("\n[+] Successfully loaded %s", configFile)
}

func SaveConfig(path string) {

	viper.SetConfigFile(path)
	if err := viper.WriteConfig(); err != nil {
		log.Fatalf("Error writing configuration file: %s", err)
	}

	color.HiMagenta("[+] Configuration saved to %s\n", path)
}
func SetupEnvironment() {
	LoadEnv()
	setupTor()
	setupLogger()
	loadCrawlerConfig()
	loadPryingConfig()
	LoadDatabase()

}
