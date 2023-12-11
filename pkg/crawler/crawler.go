package crawler

import (
	"github.com/fatih/color"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"

	"github.com/iudicium/pryingdeep/configs"
	"github.com/iudicium/pryingdeep/pkg/logger"
)

type Crawler struct {
	collector *colly.Collector
	queue     *queue.Queue
}

// NewCrawler initializes a new crawler and adds urls to the queue
func NewCrawler(torConf configs.TorConfig, crawlerConf configs.Crawler) *Crawler {
	q, _ := queue.New(
		crawlerConf.QueueThreads,
		&queue.InMemoryQueueStorage{MaxSize: crawlerConf.QueueMaxSize},
	)
	c := &Crawler{
		collector: NewCollector(crawlerConf, torConf),
		queue:     q,
	}

	c.collector.OnError(func(r *colly.Response, err error) {
		logger.Errorf("Request URL: %s failed with status code: %d Error: %s",
			r.Request.URL, r.StatusCode, err)
	})

	c.collector.OnResponse(func(response *colly.Response) {
		c.handleResponse(response, &crawlerConf)
	})

	for i, url := range crawlerConf.StartingURLS {
		logger.Infof("%v: Adding url to queue: %s", i+1, url)
		err := c.queue.AddURL(url)
		if err != nil {
			color.HiMagenta("Queue MaxSize has been reached. Exiting..")
		}
	}
	return c
}

// Crawl starts the crawling process, entrypoints are StartingURLs in the config.
// We only define the onHTML attributes here because it's easier to handle the maxSize
// Errors here, we return them and exit the application
func (c *Crawler) Crawl() error {
	var crawlErr error
	c.collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		absoluteUrl := e.Request.AbsoluteURL(url)
		err := c.queue.AddURL(absoluteUrl)
		if err != nil {
			crawlErr = err
		}
	})

	if err := c.queue.Run(c.collector); err != nil {
		return err
	}
	return crawlErr
}

// handleResponse processes the HTTP response, creates database records, and spawns goroutines
// for each corresponding module specified by the pryingConfig.json configuration
func (c *Crawler) handleResponse(response *colly.Response, options *configs.Crawler) {
	body := string(response.Body)
	url := response.Request.URL.String()
	logger.Infof("Crawling url: %s", url)
	pageId, err := ParseResponse(url, body, response)

	if err != nil {
		logger.Errorf("Something went wrong during parsing the response from: %s Err: %s ", url, err)
	}

	if options.Wordpress {
		go processWordPress(body, pageId)
	}

	if options.Email {
		go processEmail(body, pageId)
	}
	if options.Crypto {
		go processCrypto(body, pageId)
	}
	if len(options.PhoneNumbers) != 0 {
		go processPhones(body, pageId, options.PhoneNumbers)
	}

}
