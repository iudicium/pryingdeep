package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/pryingbytez/prying-deep/configs"
	"github.com/pryingbytez/prying-deep/models"
	"github.com/pryingbytez/prying-deep/pkg/crawler"
	"github.com/pryingbytez/prying-deep/pkg/logger"
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
		crawler.Crawl(cfg.TorConf, cfg.CrawlerConf, cfg.PryingConf)
	},
}

func init() {
	crawlCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", silent, "-s to disable logging and run silently")
	crawlCmd.Flags().StringVarP(&configFile, "file", "f", "configs/json/pryingConfig.json", "Configuration file path")
	crawlCmd.Flags().BoolVarP(&wordpress, "wordpress", "w", false, "Enable WordPress support")
	crawlCmd.Flags().BoolVarP(&crypto, "crypto", "c", false, "Enable crypto features")
	crawlCmd.Flags().BoolVarP(&email, "email", "e", false, "Enable email notifications")
	crawlCmd.Flags().StringSliceVarP(&phone, "phone", "p", []string{}, "Phone numbers for notifications")
}
