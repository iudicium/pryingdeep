package crawler

import (
	"errors"

	"github.com/fatih/color"
	"github.com/gocolly/colly/v2"
	"github.com/spf13/cobra"

	"github.com/iudicium/pryingdeep/configs"
	"github.com/iudicium/pryingdeep/pkg/crawler"
	"github.com/iudicium/pryingdeep/pkg/logger"
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
	setCrawlerOptions(&cfg.Crawler, cmd)
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

func setCrawlerOptions(c *configs.Crawler, cmd *cobra.Command) {
	if cmd.Flags().Changed("tor") {
		c.Tor = tor
	}
	if cmd.Flags().Changed("url-revisit") {
		c.Permissions.AllowURLRevisit = allowURLRevisit
	}
	if cmd.Flags().Changed("ignore-robots-txt") {
		c.IgnoreRobotsTxt = ignoreRobotsTxt
	}
	if cmd.Flags().Changed("debug") {
		c.Debug = debug
	}
	if cmd.Flags().Changed("wordpress") {
		c.PryingOptions.Wordpress = wordpress
	}
	if cmd.Flags().Changed("email") {
		c.PryingOptions.Email = email
	}
	if cmd.Flags().Changed("crypto") {
		c.PryingOptions.Crypto = crypto
	}

	if cmd.Flags().Changed("urls") && urls != nil {
		c.StartingURLS = urls
	}
	if cmd.Flags().Changed("max-depth") && maxDepth != 0 {
		c.MaxDepth = maxDepth
	}
	if cmd.Flags().Changed("bod-size") && maxBodySize != 0 {
		c.MaxBodySize = maxBodySize
	}
	if cmd.Flags().Changed("cache-dir") && cacheDir != "" {
		c.CacheDir = cacheDir
	}

	if cmd.Flags().Changed("queue-threads") && queueThreads != 0 {
		c.QueueThreads = queueThreads
	}

	if cmd.Flags().Changed("queue-max-size") && queueMaxSize != 0 {
		c.QueueMaxSize = queueMaxSize
	}

	if cmd.Flags().Changed("allowed-domains") && allowedDomains != nil {
		c.Permissions.AllowedDomains = allowedDomains
	}

	if cmd.Flags().Changed("disallowed-url-filters") && disallowedURLFilters != nil {
		c.Permissions.DisallowedURLFilters = disallowedURLFilters
	}
	if cmd.Flags().Changed("disallowed-domains") && disallowedDomains != nil {
		c.Permissions.DisallowedDomains = disallowedDomains
	}
	if cmd.Flags().Changed("url-filters") && urlFilters != nil {
		c.Permissions.URLFilters = urlFilters
	}

	if cmd.Flags().Changed("limit-delay") && delay != 0 {
		c.LimitRule.Delay = delay
	}

	if cmd.Flags().Changed("limit-random-delay") && randomDelay != 0 {
		c.LimitRule.RandomDelay = randomDelay
	}
	if cmd.Flags().Changed("phone") && phone != nil {
		c.PryingOptions.PhoneNumbers = phone
	}
}
