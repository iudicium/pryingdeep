package cmd

import (
	"errors"

	"github.com/fatih/color"
	"github.com/gocolly/colly/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/pryingbytez/pryingdeep/configs"
	"github.com/pryingbytez/pryingdeep/models"
	"github.com/pryingbytez/pryingdeep/pkg/crawler"
	"github.com/pryingbytez/pryingdeep/pkg/logger"
)

var (
	silent     = false
	configFile string
	wordpress  bool
	crypto     bool
	email      bool
	phone      []string
)
var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Start the crawling process",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("WordPress", wordpress)
		viper.Set("Crypto", crypto)
		viper.Set("Email", email)
		viper.Set("PhoneNumbers", phone)

		configs.SaveConfig(configFile)

		configs.SetupEnvironment()
		cfg := configs.GetConfig()
		logger.InitLogger(silent)
		defer logger.Logger.Sync()

		models.SetupDatabase(cfg.DbConf.DbURL)
		crawler := crawler.NewCrawler(cfg.TorConf, cfg.CrawlerConf, cfg.PryingConf)
		if err := crawler.Crawl(); err != nil {
			if errors.Is(err, colly.ErrQueueFull) {
				color.HiRed("\nQueue max size has been reached! Exiting.")
			} else {
				logger.Errorf("Crawl error: %s", err)
			}
		}

	},
}

func init() {
	crawlCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", silent, "-s to disable logging and run silently")
	crawlCmd.Flags().StringVarP(&configFile, "file", "f", "configs/json/pryingConfig.json", "Configuration file path")
	crawlCmd.Flags().BoolVarP(&wordpress, "wordpress", "w", false, "Enable WordPress support")
	crawlCmd.Flags().BoolVarP(&crypto, "crypto", "c", false, "Enable crypto features")
	crawlCmd.Flags().BoolVarP(&email, "email", "e", false, "Enable email notifications")
	crawlCmd.Flags().StringSliceVarP(&phone, "phone", "p", []string{}, "List of countries. RU,NL,DE,GB,US. You can specify multiple or just one.")
}
