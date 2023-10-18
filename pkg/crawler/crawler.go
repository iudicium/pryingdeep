package crawler

import (
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
	"github.com/r00tk3y/prying-deep/configs"
	"github.com/r00tk3y/prying-deep/pkg/logger"
	"strings"
)

func Crawl(torConf configs.TorConfig, crawlerConf configs.CollyConfig) {
	c := NewCollector(crawlerConf, torConf)

	q, _ := queue.New(
		crawlerConf.QueueThreads,
		&queue.InMemoryQueueStorage{MaxSize: crawlerConf.QueueMaxSize},
	)

	//TODO: add  error handling to separate file
	c.OnError(func(_ *colly.Response, err error) {
		//TODO: This approach works for now but will need to a create custom error types in the future
		if strings.Contains(err.Error(), "socks connect tcp 127.0.0.1") {
			logger.Errorf("Website does not exist!")
		} else {
			logger.Errorf("Something went wrong: %s", err)
		}
	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		e.Request.Visit(url)
	})

	c.OnResponse(func(response *colly.Response) {
		HandleResponse(response)
	})

	for i, url := range crawlerConf.StartingURLS {
		logger.Infof("%v: Adding url to queue: %s", i, url)
		q.AddURL(url)
	}

	err := q.Run(c)
	if err != nil {
		logger.Errorf("que run err %s", err)
	}
}
