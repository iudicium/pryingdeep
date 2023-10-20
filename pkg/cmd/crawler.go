package cmd

import (
	"fmt"
	"github.com/r00tk3y/prying-deep/configs"
	"github.com/r00tk3y/prying-deep/models"
	"github.com/r00tk3y/prying-deep/pkg/crawler"
	"github.com/r00tk3y/prying-deep/pkg/logger"
	"github.com/spf13/cobra"
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
	RootCmd.AddCommand(crawlCmd)
}
