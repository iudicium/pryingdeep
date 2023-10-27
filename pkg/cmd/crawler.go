package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/pryingbytez/prying-deep/configs"
	"github.com/pryingbytez/prying-deep/models"
	"github.com/pryingbytez/prying-deep/pkg/crawler"
	"github.com/pryingbytez/prying-deep/pkg/logger"
)

var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Start the crawling process",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(art())
		configs.SetupEnvironment()
		cfg := configs.GetConfig()
		logger.InitLogger()
		defer logger.Logger.Sync()
		models.SetupDatabase(cfg.DbConf.DbURL)
		crawler.Crawl(cfg.TorConf, cfg.CrawlerConf, cfg.PryingConf)
	},
}

func init() {
	rootCmd.AddCommand(crawlCmd)
}
