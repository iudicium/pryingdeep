package crawler

import (
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
	"github.com/r00tk3y/prying-deep/configs"
	"github.com/r00tk3y/prying-deep/pkg/logger"
)

func Crawl(torConf configs.TorConfig, crawlerConf configs.CollyConfig) {
	c := NewCollector(crawlerConf, torConf)

	q, _ := queue.New(
		crawlerConf.QueueThreads,
		&queue.InMemoryQueueStorage{MaxSize: crawlerConf.QueueMaxSize},
	)

	// c.OnError(func(_ *colly.Response, err error) {
	//    logger.Error("Something went wrong: ", zap.Error(err))
	// })
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		ParseAndAddURL(q, e)
	})

	c.OnResponse(func(response *colly.Response) {
		HandleResponse(response)
	})

	q.AddURL(crawlerConf.StartingURL)

	err := q.Run(c)
	if err != nil {
		logger.Errorf("que run err %s", err)
	}
}
