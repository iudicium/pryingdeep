package main


import (
	"github.com/r00tk3y/prying-deep/pkg/crawler"
	"github.com/r00tk3y/prying-deep/pkg/logger"
	"github.com/r00tk3y/prying-deep/configs"
	
)


// var torProxy string = "socks5://127.0.0.1:9050"


func main() {
	logger  := logger.NewLogger()
	defer logger.Sync()

	environmentVaribles := configs.SetupEnvironment()

	logger.Info("Starting the crawl process")
	crawler.Crawl("http://paavlaytlfsqyvkg3yqj7hflfg5jw2jdg2fgkza5ruf6lplwseeqtvyd.onion/", environmentVaribles.TorConf)
	for {
							
	}
}


// #https://stackoverflow.com/questions/63636649/how-do-i-connect-a-docker-container-to-tor-proxy-on-local-machine