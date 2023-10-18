package main

import (
	"github.com/r00tk3y/prying-deep/configs"
	"github.com/r00tk3y/prying-deep/models"
	//"github.com/r00tk3y/prying-deep/pkg/cli"
	"github.com/r00tk3y/prying-deep/pkg/crawler"
	"github.com/r00tk3y/prying-deep/pkg/logger"
)

func main() {
	configs.SetupEnvironment()
	cfg := configs.GetConfig()
	logger.InitLogger()
	defer logger.Logger.Sync()

	//TODO add sitemap, wordpress, nginx searching for the tool
	//TODO, enhance regexp for USA phone numbers it's giving too much false positives, maybe only look for internaitonal phone numbers
	logger.Info("Configuring Database and Running migrations")
	models.SetupDatabase(cfg.DbConf.DbURL)

	logger.Info("Starting the crawl process")

	crawler.Crawl(cfg.TorConf, cfg.CrawlerConf)

}

//view-source:http://xjfbpuj56rdazx4iolylxplbvyft2onuerjeimlcqwaihp3s6r4xebqd.onion/chatgpt-web-crawler/comment-page-1/
