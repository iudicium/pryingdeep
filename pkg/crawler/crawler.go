package crawler

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
	"go.uber.org/zap"

	"github.com/r00tk3y/prying-deep/configs"
	"github.com/r00tk3y/prying-deep/pkg/logger"
	"github.com/r00tk3y/prying-deep/pkg/parsers"
	"github.com/r00tk3y/prying-deep/pkg/pryingtools/email"
	"github.com/r00tk3y/prying-deep/pkg/pryingtools/wordpress"
	"github.com/r00tk3y/prying-deep/pkg/utils"
)

func Crawl(urlToCrawl string, socks5conf configs.TorConfig, maxDepth int, ua string) {

	torProxy := fmt.Sprintf("socks5://%s:%s", socks5conf.Host, socks5conf.Port)

	c := colly.NewCollector(
		colly.UserAgent(ua),
		colly.IgnoreRobotsTxt(),
		colly.TraceHTTP(),
		colly.MaxDepth(maxDepth),
	)
	rp, err := proxy.RoundRobinProxySwitcher(torProxy)
	if err != nil {
		logger.Fatal(err.Error())
	}
	c.SetProxyFunc(rp)

	checkTorConnection, err := utils.CheckIfTorConnectionExists(torProxy)
	if !checkTorConnection {
		logger.Fatal("Killing session. Can't connect to Tor.\nError:  ", zap.Error(err))
	}

	// c.OnError(func(_ *colly.Response, err error) {
	//    logger.Error("Something went wrong: ", zap.Error(err))
	// })
	var title string
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		//fmt.Println(e.Attr("href"))
		e.Request.Visit(e.Attr("href"))
	})
	c.OnHTML("head title", func(e *colly.HTMLElement) {
		title = e.Text
	})

	c.OnResponse(func(response *colly.Response) {
		logger.Infof("Title: ", title)

		body := string(response.Body)
		url := response.Request.URL.String()
		logger.Infof("Crawling url:", url)

		wordpressMatches, _ := wordpress.FindWordpressPatterns(body, url)
		logger.Infof("Wordpress matches", wordpressMatches)

		emailMatches := email.FindEmail(body)
		logger.Infof("Email matches: ", emailMatches)
		_, err := parsers.ParseResponse(response)
		if err != nil {
			logger.Errorf("Something went wrong during parsing the response from: ", url)
		}
	})
	err = c.Visit(urlToCrawl)
	if err != nil {
		logger.Debug("Something went wrong", zap.Error(err))
	}
}
