package crawler

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
    "github.com/gocolly/colly/v2/proxy"

	"github.com/r00tk3y/prying-deep/configs"
)




func Crawl(url string, socks5conf configs.Socks5Config) {

	torProxy := fmt.Sprintf("socks5://%s:%s", socks5conf.Host, socks5conf.Port)
	checkTor := "https://check.torproject.org/"

	c := colly.NewCollector(colly.AllowURLRevisit())

	rp, err := proxy.RoundRobinProxySwitcher(torProxy)
	if err != nil {
		log.Fatal(err)
	}

	c.SetProxyFunc(rp)

	c.OnError(func(_ *colly.Response, err error) {
       log.Println("Something went wrong:", err)
    })
	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println(e.Request.URL)
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.Body)
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(checkTor)
	
}