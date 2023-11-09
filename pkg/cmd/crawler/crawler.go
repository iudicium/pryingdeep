package crawler

import (
	"errors"

	"github.com/fatih/color"
	"github.com/gocolly/colly/v2"
	"github.com/spf13/cobra"

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
	cliConfig *configs.CLIConfig
)
var CrawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Start the crawling process",
	Run: func(cmd *cobra.Command, args []string) {

		cliConfig = configs.NewCLIConfig()
		configs.SetupEnvironment()

		cfg := configs.GetConfig()
		if cliConfig.SaveConfig {
			cliConfig.StoreConfig()
		}
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
	CrawlCmd.Flags().BoolVarP(&wordpress, "wordpress", "w", wordpress, "Enable WordPress support")
	CrawlCmd.Flags().BoolVarP(&crypto, "crypto", "k", crypto, "Enable crypto features")
	CrawlCmd.Flags().BoolVarP(&email, "email", "e", email, "Enable email search")

	CrawlCmd.Flags().BoolVarP(&tor, "tor", "t", true, "Turn off connecting to tor and crawl with your IP on the clearweb. -t=false")
	CrawlCmd.Flags().StringSliceVarP(&phone, "phone", "p", phone, "List of countries. RU,NL,DE,GB,US. You can specify multiple or just one.")

}
