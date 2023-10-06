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

	logger.Info("Configuring Database and Running migrations")
	models.SetupDatabase(&cfg.DbConf)
	logger.Info("Starting the crawl process")

	// Allowed domains are specified splitting by ,
	crawler.Crawl("http://paavlaytlfsqyvkg3yqj7hflfg5jw2jdg2fgkza5ruf6lplwseeqtvyd.onion/", cfg.TorConf, 5)

}

// #https://stackoverflow.com/questions/63636649/how-do-i-connect-a-docker-container-to-tor-proxy-on-local-machine
