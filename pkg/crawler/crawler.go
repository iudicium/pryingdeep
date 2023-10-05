package crawler

import (
	"fmt"

	"go.uber.org/zap"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"

	"github.com/r00tk3y/prying-deep/pkg/parsers"
	"github.com/r00tk3y/prying-deep/configs"
	"github.com/r00tk3y/prying-deep/pkg/logger"

)



func Crawl(urlToCrawl string, socks5conf configs.Socks5Config, maxDepth int) {
	logger := logger.NewLogger()

	
	torProxy := fmt.Sprintf("socks5://%s:%s", socks5conf.Host, socks5conf.Port)

	c := colly.NewCollector(
		colly.IgnoreRobotsTxt(),
		colly.TraceHTTP(),
		colly.MaxDepth(maxDepth),
	
	)
	rp, err := proxy.RoundRobinProxySwitcher(torProxy)
	if err != nil {
		logger.Fatal(err.Error())
	}
	c.SetProxyFunc(rp)
	
	checkTorConnection, err := checkIfTorConnectionExists(torProxy)
	if !checkTorConnection {
		logger.Fatal("Killing session. Can't connect to Tor.\nError:  ",  zap.Error(err))
	}

	// c.OnError(func(_ *colly.Response, err error) {
    //    logger.Error("Something went wrong: ", zap.Error(err))
    // })
	// // Find and visit all links
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	fmt.Println(e)
	// 	e.Request.Visit(e.Attr("href"))
	// })

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.Headers)
		fmt.Println("Visiting", r.URL)
		parsers.ParseRequest(r, logger)
	})

	
	
	c.OnResponse(func(response *colly.Response) {
		logger.Info("Request Method", zap.Any("Trace callback", response.Trace))

		fmt.Println("Visited:", response.Request.URL)
		fmt.Println("Response Status Code:", response.StatusCode)
		fmt.Println("Response Body:", string(response.Body))
	})
	c.Visit(urlToCrawl)
}