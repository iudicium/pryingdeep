package main


import (
	"fmt"


	"github.com/r00tk3y/prying-deep/pkg/crawler"
	"github.com/r00tk3y/prying-deep/configs"
	
)


// var torProxy string = "socks5://127.0.0.1:9050"


func main() {

	environmentVaribles := configs.SetupEnvironment()

	fmt.Println("Starting")
	crawler.Crawl("https://google.com/", environmentVaribles.TorConf)
	for {
							
	}
}


// #https://stackoverflow.com/questions/63636649/how-do-i-connect-a-docker-container-to-tor-proxy-on-local-machine