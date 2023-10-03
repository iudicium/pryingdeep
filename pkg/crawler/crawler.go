package crawler

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gocolly/colly/v2"

	"github.com/r00tk3y/prying-deep/configs"
)


const checkTor string = "https://check.torproject.org/"

func Crawl(urlToCrawl string, socks5conf configs.Socks5Config) {

	torProxy := fmt.Sprintf("socks5://%s:%s", socks5conf.Host, socks5conf.Port)
	

	torProxyUrl, err := url.Parse(torProxy)
	if err != nil {
		log.Fatal("Error parsing Tor proxy URL:", torProxy, ".", err)
	}
	
	torTransport := &http.Transport{Proxy: http.ProxyURL(torProxyUrl)}
	

	c := colly.NewCollector()
	c.WithTransport(torTransport)	


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