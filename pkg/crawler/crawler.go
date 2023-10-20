package crawler

import (
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
	"github.com/r00tk3y/prying-deep/configs"
	"github.com/r00tk3y/prying-deep/pkg/logger"
)

func Crawl(torConf configs.TorConfig, crawlerConf configs.CollyConfig, pryingConf configs.PryingConfig) {
	c := NewCollector(crawlerConf, torConf)
	q, _ := queue.New(
		crawlerConf.QueueThreads,
		&queue.InMemoryQueueStorage{MaxSize: crawlerConf.QueueMaxSize},
	)

	//TODO: add  error handling to separate file
	c.OnError(func(r *colly.Response, err error) {
		logger.Errorf("Request URL: %s failed with status code: %d Error: %s",
			r.Request.URL, r.StatusCode, err)
	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		e.Request.Visit(url)
	})
	c.OnResponse(func(response *colly.Response) {
		HandleResponse(response, &pryingConf)
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
