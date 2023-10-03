package crawler

import (
	"fmt"
	"log"


	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"

	"github.com/r00tk3y/prying-deep/configs"
)



func Crawl(urlToCrawl string, socks5conf configs.Socks5Config) {

	torProxy := fmt.Sprintf("socks5://%s:%s", socks5conf.Host, socks5conf.Port)

	c := colly.NewCollector()
	rp, err := proxy.RoundRobinProxySwitcher(torProxy)
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)
	
	checkTorConnection, err := checkIfTorConnectionExists(torProxy)
	if !checkTorConnection {
 	   log.Fatal("Killing session. Can't connect to Tor. Error: ", err )
	}

	// c.OnError(func(_ *colly.Response, err error) {
    //    log.Println("Something went wrong:", err)
    // })
	// // Find and visit all links
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	fmt.Println(e)
	// 	e.Request.Visit(e.Attr("href"))
	// })

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println(r.Headers)
	// 	fmt.Println("Visiting", r.URL)
	// })

	
	c.OnResponse(func(response *colly.Response) {

		fmt.Println("Visited:", response.Request.URL)
		fmt.Println("Response Status Code:", response.StatusCode)
		fmt.Println("Response Body:", string(response.Body))
	})
	c.Visit(urlToCrawl)
}