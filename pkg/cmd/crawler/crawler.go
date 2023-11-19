package crawler

import (
	"errors"

	"github.com/fatih/color"
	"github.com/gocolly/colly/v2"
	"github.com/spf13/cobra"

	"github.com/pryingbytez/pryingdeep/configs"
	"github.com/pryingbytez/pryingdeep/pkg/crawler"
	"github.com/pryingbytez/pryingdeep/pkg/logger"
)

var (
	//Crawler options
	urls                 []string
	tor                  bool
	userAgent            string
	maxDepth             int
	maxBodySize          int
	cacheDir             string
	ignoreRobotsTxt      bool
	debug                bool
	queueThreads         int
	queueMaxSize         int
	allowedDomains       []string
	disallowedDomains    []string
	disallowedURLFilters []string
	urlFilters           []string
	allowURLRevisit      bool
	delay                int
	randomDelay          int
	// Prying options - (more to come!)
	wordpress bool
	crypto    bool
	email     bool
	phone     []string
)

var CrawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Start the crawling process",
	Run:   Crawl,
}

func Crawl(cmd *cobra.Command, args []string) {
	configs.SetupEnvironment()
	cfg := configs.GetConfig()
	setCrawlerOptions(&cfg.Crawler)
	newCrawler := crawler.NewCrawler(cfg.Tor, cfg.Crawler)
	if err := newCrawler.Crawl(); err != nil {
		if errors.Is(err, colly.ErrQueueFull) {
			color.HiRed("\nQueue max size has been reached! Exiting.")
		} else {
			logger.Errorf("Crawl error: %s", err)
		}
	}

}

func init() {
	CrawlCmd.Flags().StringSliceVarP(&urls, "urls", "u", nil, "Entry point URLs")
	CrawlCmd.Flags().BoolVarP(&tor, "tor", "t", false, "Turn on/off connecting to Tor.")
	CrawlCmd.Flags().StringVar(&userAgent, "user-agent", userAgent, "Specify any user agents for the crawler to use.")
	CrawlCmd.Flags().IntVar(&maxDepth, "max-depth", 0, "Maximum recursion depth")
	CrawlCmd.Flags().IntVar(&maxBodySize, "body-size", 0, "Max body size in bytes (0 for unlimited)")
	CrawlCmd.Flags().StringVar(&cacheDir, "cache-dir", "", "Cache directory")
	CrawlCmd.Flags().BoolVar(&ignoreRobotsTxt, "ignore-robots-txt", false, "Ignore robots.txt")
	CrawlCmd.Flags().BoolVar(&debug, "debug", false, "Enable debug mode")
	CrawlCmd.Flags().IntVar(&queueThreads, "queue-threads", 4, "Number of queue threads")
	CrawlCmd.Flags().IntVar(&queueMaxSize, "queue-max-size", 50000, "Queue max size")

	CrawlCmd.Flags().StringSliceVar(&allowedDomains, "allowed-domains", nil, "Allowed domains")
	CrawlCmd.Flags().StringSliceVar(&disallowedDomains, "disallowed-domains", nil, "Disallowed domains")
	CrawlCmd.Flags().StringSliceVar(&disallowedURLFilters, "disallowed-url-filters", nil, "Disallowed URL filters")
	CrawlCmd.Flags().StringSliceVar(&urlFilters, "url-filters", nil, "URL filters")
	CrawlCmd.Flags().BoolVar(&allowURLRevisit, "url-revisit", false, "Allow URL revisit")

	CrawlCmd.Flags().IntVar(&delay, "limit-delay", 0, "Limit delay")
	CrawlCmd.Flags().IntVar(&randomDelay, "limit-random-delay", 0, "Limit random delay")

	CrawlCmd.Flags().BoolVarP(&wordpress, "wordpress", "w", false, "Enable WordPress support")
	CrawlCmd.Flags().BoolVarP(&crypto, "crypto", "k", false, "Enable crypto features")
	CrawlCmd.Flags().BoolVarP(&email, "email", "e", false, "Enable email search")
	CrawlCmd.Flags().StringSliceVarP(&phone, "phone", "p", []string{}, "List of countries. RU,NL,DE,GB,US. You can specify multiple or just one.")

	cli := configs.NewCLIConfig()
	CrawlCmd.Flags().VisitAll(cli.ConfigureViper("crawler"))
}

func setCrawlerOptions(c *configs.Crawler) {
	//Only bool flags here, no need for if statements.
	c.Tor = tor
	c.Permissions.AllowURLRevisit = allowURLRevisit
	c.IgnoreRobotsTxt = ignoreRobotsTxt
	c.Debug = debug
	c.PryingOptions.Wordpress = wordpress
	c.PryingOptions.Email = email
	c.PryingOptions.Crypto = crypto

	if urls != nil {
		c.StartingURLS = urls
	}
	if maxDepth != 0 {
		c.MaxDepth = maxDepth
	}
	if maxBodySize != 0 {
		c.MaxBodySize = maxBodySize
	}
	if cacheDir != "" {
		c.CacheDir = cacheDir
	}

	if queueThreads != 0 {
		c.QueueThreads = queueThreads
	}
	if queueMaxSize != 0 {
		c.QueueMaxSize = queueMaxSize
	}
	if allowedDomains != nil {
		c.Permissions.AllowedDomains = allowedDomains
	}
	if disallowedURLFilters != nil {
		c.Permissions.DisallowedURLFilters = disallowedURLFilters
	}
	if disallowedDomains != nil {
		c.Permissions.DisallowedDomains = disallowedDomains
	}
	if urlFilters != nil {
		c.Permissions.URLFilters = urlFilters
	}

	if delay != 0 {
		c.LimitRule.Delay = delay
	}
	if randomDelay != 0 {
		c.LimitRule.RandomDelay = randomDelay
	}
	if phone != nil {
		c.PryingOptions.PhoneNumbers = phone
	}
}
