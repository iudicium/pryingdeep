package configs

import (
	"log"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"

	"github.com/iudicium/pryingdeep/pkg/fsutils"
)

// Configuration holds different components for easy access.
type Configuration struct {
	Tor      TorConfig
	DB       Database
	Logger   LoggerConfig
	Crawler  Crawler
	Exporter Exporter
}

var cfg Configuration

func GetConfig() *Configuration {
	return &cfg
}

func loadConfig(key string, config interface{}) {
	if err := viper.UnmarshalKey(key, config); err != nil {
		log.Fatal(err)
	}

}

func Save(ignoredkeys ...string) error {
	file := viper.ConfigFileUsed()
	if len(file) == 0 {
		file = "./pryingdeep.yaml"
	}

	configMap := viper.AllSettings()
	for _, key := range ignoredkeys {
		delete(configMap, key)
	}
	content, err := yaml.Marshal(configMap)
	if err != nil {
		return err
	}

	fsutils.WriteTextFile(file, string(content))

	return nil
}

func SetupEnvironment() {
	setupTor()
	setupLogger()
	loadCrawlerConfig()
	LoadDatabase()
}
