package main

import (
	"context"
	"database/sql"

	"go.uber.org/zap"
	"github.com/r00tk3y/prying-deep/configs"
	"github.com/r00tk3y/prying-deep/pkg/crawler"
	"github.com/r00tk3y/prying-deep/pkg/logger"
	"github.com/r00tk3y/prying-deep/internal/database"
)

// var torProxy string = "socks5://127.0.0.1:9050"



func main() {
	logger  := logger.NewLogger()
	defer logger.Sync()

	EnvConfigs := configs.SetupEnvironment()


	ctx := context.Background()

	db, err := sql.Open("postgres", EnvConfigs.DbConf.DbURL)
	if err != nil {
		logger.Fatal("Error opening the database connection", zap.Error(err))
	}
	dbQueries := 
	logger.Info("Starting the crawl process")
	// Allowed domains are specified splitting by ,
	crawler.Crawl("http://paavlaytlfsqyvkg3yqj7hflfg5jw2jdg2fgkza5ruf6lplwseeqtvyd.onion/", EnvConfigs.TorConf, 5)
	
}


// #https://stackoverflow.com/questions/63636649/how-do-i-connect-a-docker-container-to-tor-proxy-on-local-machine