package main

import (
	"github.com/r00tk3y/prying-deep/configs"
	"github.com/r00tk3y/prying-deep/models"
	"github.com/r00tk3y/prying-deep/pkg/crawler"
	"github.com/r00tk3y/prying-deep/pkg/logger"
)

func main() {
	logger.InitLogger()
	defer logger.Logger.Sync()
	configs.SetupEnvironment()
	cfg := configs.GetConfig()
	//TODO add sitemap, wordpress, nginx searching for the tool

	logger.Info("Configuring Database and Running migrations")
	models.SetupDatabase(cfg.DbConf.DbURL)

	logger.Info("Starting the crawl process")
	crawler.Crawl(cfg.TorConf, cfg.CrawlerConf)

}

//TODO: ADD a testing database
// #https://stackoverflow.com/questions/63636649/how-do-i-connect-a-docker-container-to-tor-proxy-on-local-machine
//view-source:http://xjfbpuj56rdazx4iolylxplbvyft2onuerjeimlcqwaihp3s6r4xebqd.onion/chatgpt-web-crawler/comment-page-1/
