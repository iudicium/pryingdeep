package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/pryingbytez/pryingdeep/configs"
	"github.com/pryingbytez/pryingdeep/models"
	"github.com/pryingbytez/pryingdeep/pkg/cmd/crawler"
	"github.com/pryingbytez/pryingdeep/pkg/cmd/exporter"
	"github.com/pryingbytez/pryingdeep/pkg/logger"
)

var (
	saveConfig = false
	silent     = false
	path       = "pryingdeep.yaml"
)

var rootCmd = &cobra.Command{
	Use:   "pryingdeep",
	Short: "Pryingdeep is a dark web osint intelligence tool.",
	Long: `Pryingdeep specializes in collecting information about dark-web/clearnet websites.
		This tool was specifically built to extract as much information as possible from a .onion website`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := initializeConfig(cmd, args)
		configs.LoadDatabase()
		logger.InitLogger(silent)
		defer logger.Logger.Sync()

		cfg := configs.GetConfig()
		models.SetupDatabase(cfg.DB.URL)
		return err
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", silent, "-s to disable logging and run silently")
	rootCmd.PersistentFlags().BoolVarP(&saveConfig, "save-config", "z", saveConfig, "Save chosen options to your .yaml configuration")
	rootCmd.PersistentFlags().StringVarP(&path, "config", "c", path, "Path to the .yaml configuration.")
	rootCmd.AddCommand(exporter.ExporterCMD)
	rootCmd.AddCommand(crawler.CrawlCmd)
	viper.BindPFlags(rootCmd.PersistentFlags())

}

func initializeConfig(cmd *cobra.Command, args []string) error {
	viper.SetConfigType("yaml")
	if path != "" {
		viper.SetConfigFile(path)
		fmt.Println("test", path)
	}

	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.pryingdeep")

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			color.Red("Config file has not been found. " +
				"Please download it from our github and place it inside $HOME/.pryingdeep")
		}
		os.Exit(1)

	}

	viper.SetEnvPrefix("DEEP")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	bindFlags(cmd)

	cliConfig := configs.NewCLIConfig()
	if saveConfig {
		cliConfig.StoreConfig()
	}

	return nil

}

// Bind each cobra flag to its associated viper configuration (config file and environment variable)
func bindFlags(cmd *cobra.Command) {
	call := func(f *pflag.Flag) {
		configName := f.Name
		if !f.Changed && viper.IsSet(configName) {
			val := viper.Get(configName)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))

		}

	}
	cmd.Flags().VisitAll(call)
}
