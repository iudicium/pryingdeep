package helpers

import (
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/viper"

	"github.com/iudicium/pryingdeep/configs"
	"github.com/iudicium/pryingdeep/models"
	"github.com/iudicium/pryingdeep/pkg/logger"
)

func InitTestConfig() error {
	viper.SetConfigName("pryingdeep")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.pryingdeep")

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			color.Red("Config file has not been found. " +
				"Please download it from our github and place it inside $HOME/.pryingdeep")
			os.Exit(1)
		} else {
			return err
		}

	}

	configs.LoadDatabase()
	logger.InitLogger(false)
	defer logger.Logger.Sync()

	cfg := configs.GetConfig().DB

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.TestName)
	models.SetupDatabase(dbURL)
	configs.SetupEnvironment()
	return nil
}
