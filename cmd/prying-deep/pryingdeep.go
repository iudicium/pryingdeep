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
	//TODO add .env support for user agent
	//TODO add sitemap, wordpress, nginx searching for the tool
	//TODO: Remove Response Model, unnecessary fields

	ua := "Mozilla/5.0 (X11; Linux x86_64; rv:102.0) Gecko/20100101 Firefox/102.0"
	logger.Info("Configuring Database and Running migrations")
	models.SetupDatabase(&cfg.DbConf)
	logger.Info("Starting the crawl process")

	// Allowed domains are specified splitting by ,
	crawler.Crawl("http://xjfbpuj56rdazx4iolylxplbvyft2onuerjeimlcqwaihp3s6r4xebqd.onion/", cfg.TorConf, 2, ua)

}

// #https://stackoverflow.com/questions/63636649/how-do-i-connect-a-docker-container-to-tor-proxy-on-local-machine
//view-source:http://xjfbpuj56rdazx4iolylxplbvyft2onuerjeimlcqwaihp3s6r4xebqd.onion/chatgpt-web-crawler/comment-page-1/
