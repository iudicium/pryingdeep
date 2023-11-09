package config

import (
	"errors"
	"fmt"

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
	wordpress bool
	crypto    bool
	email     bool
	tor       bool
	phone     []string
)
var CrawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Start the crawling process",
	Run: func(cmd *cobra.Command, args []string) {

		viperConfig := viper.GetViper()
		fmt.Println(viperConfig)
		cliConfig := configs.NewCLIConfig(viperConfig)

		configs.SetupEnvironment()
		cfg := configs.GetConfig()
		logger.InitLogger(cliConfig.Silent)
		defer logger.Logger.Sync()

		models.SetupDatabase(cfg.DB.URL)
		crawler := crawler.NewCrawler(cfg.Tor, cfg.Crawler, cfg.PryingOptions)
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
	CrawlCmd.Flags().BoolVarP(&wordpress, "wordpress", "w", false, "Enable WordPress support")
	CrawlCmd.Flags().BoolVarP(&crypto, "crypto", "k", false, "Enable crypto features")
	CrawlCmd.Flags().BoolVarP(&email, "email", "e", false, "Enable email search")

	CrawlCmd.Flags().BoolVarP(&tor, "tor", "t", true, "Turn off connecting to tor and crawl with your IP on the clearweb. -t=false")
	CrawlCmd.Flags().StringSliceVarP(&phone, "phone", "p", []string{}, "List of countries. RU,NL,DE,GB,US. You can specify multiple or just one.")

}
