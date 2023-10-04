package crawler

import (
	"fmt"
	"reflect"
	"go.uber.org/zap"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"

	"github.com/r00tk3y/prying-deep/configs"
	"github.com/r00tk3y/prying-deep/pkg/logger"

)


func ParseRequest(request *colly.Request, logger *zap.Logger) {
	test := fmt.Sprintf("%v", reflect.TypeOf(request.Headers))
	logger.Debug(test)

	logger.Debug(test)
	logger.Info("Request Method", zap.String("Method", request.Method))
	logger.Info("Request URL", zap.String("URL", request.URL.String()))
	logger.Info("Request Headers", zap.Any("Headers", request.Headers))
	logger.Info("User-Agent", zap.String("UserAgent", request.Headers.Get("User-Agent")))
	logger.Info("Referer", zap.String("Referer", request.Headers.Get("Referer")))
	logger.Info("Cookies", zap.String("Cookies", request.Headers.Get("Cookie")))
}
func Crawl(urlToCrawl string, socks5conf configs.Socks5Config) {
	logger := logger.NewLogger()

	
	torProxy := fmt.Sprintf("socks5://%s:%s", socks5conf.Host, socks5conf.Port)

	c := colly.NewCollector()
	rp, err := proxy.RoundRobinProxySwitcher(torProxy)
	if err != nil {
		logger.Fatal(err.Error())
	}
	c.SetProxyFunc(rp)
	
	checkTorConnection, err := checkIfTorConnectionExists(torProxy)
	if !checkTorConnection {
		logger.Fatal("Killing session. Can't connect to Tor.\nError:  ",  zap.Error(err))
	}

	c.OnError(func(_ *colly.Response, err error) {
       logger.Error("Something went wrong: ", zap.Error(err))
    })
	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println(e)
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.Headers)
		fmt.Println("Visiting", r.URL)
		ParseRequest(r, logger)
	})

	
	c.OnResponse(func(response *colly.Response) {

		fmt.Println("Visited:", response.Request.URL)
		fmt.Println("Response Status Code:", response.StatusCode)
		fmt.Println("Response Body:", string(response.Body))
	})
	c.Visit(urlToCrawl)
}