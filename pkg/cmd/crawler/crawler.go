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

var CrawlCMD = &cobra.Command{
	Use:   "crawl",
	Short: "Start the crawling process",
	Run: func(cmd *cobra.Command, args []string) {
		setupCrawlerConfig(cmd)
		Crawl()
	},
}

func init() {
	initCrawler(CrawlCMD)
}

func setupCrawlerConfig(cmd *cobra.Command) *configs.Configuration {
	configs.SetupEnvironment()
	cfg = configs.GetConfig()
	setCrawlerOptions(&cfg.Crawler, cmd)

	return cfg
}

func handleCrawlError(err error) {
	if errors.Is(err, colly.ErrQueueFull) {
		color.HiRed("\nQueue max size has been reached! Exiting.")
	} else {
		logger.Errorf("Crawl error: %s", err)
	}
}

func Crawl() {
	newCrawler := crawler.NewCrawler(cfg.Tor, cfg.Crawler)
	if err := newCrawler.Crawl(); err != nil {
		handleCrawlError(err)
	}
}

func initCrawler(cmd *cobra.Command) {
	cmd.PersistentFlags().StringSliceVarP(&urls, "urls", "u", nil, "Entry point URLs")
	cmd.PersistentFlags().BoolVarP(&tor, "tor", "t", false, "Turn on/off connecting to Tor.")
	cmd.PersistentFlags().StringVar(&userAgent, "user-agent", userAgent, "Specify any user agents for the crawler to use.")

	cmd.PersistentFlags().IntVar(&maxDepth, "max-depth", 0, "Maximum recursion depth")
	cmd.PersistentFlags().IntVar(&maxBodySize, "body-size", 0, "Max body size in bytes (0 for unlimited)")
	cmd.PersistentFlags().StringVar(&cacheDir, "cache-dir", "", "Cache directory")
	cmd.PersistentFlags().BoolVar(&ignoreRobotsTxt, "ignore-robots-txt", false, "Ignore robots.txt")
	cmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug mode")
	cmd.PersistentFlags().IntVar(&queueThreads, "queue-threads", 4, "Number of queue threads")
	cmd.PersistentFlags().IntVar(&queueMaxSize, "queue-max-size", 50000, "Queue max size")

	cmd.PersistentFlags().StringSliceVar(&allowedDomains, "allowed-domains", nil, "Allowed domains")
	cmd.PersistentFlags().StringSliceVar(&disallowedDomains, "disallowed-domains", nil, "Disallowed domains")
	cmd.PersistentFlags().StringSliceVar(&disallowedURLFilters, "disallowed-url-filters", nil, "Disallowed URL filters")
	cmd.PersistentFlags().StringSliceVar(&urlFilters, "url-filters", nil, "URL filters")
	cmd.PersistentFlags().BoolVar(&allowURLRevisit, "url-revisit", false, "Allow URL revisit")

	cmd.PersistentFlags().IntVar(&delay, "limit-delay", 0, "Limit delay")
	cmd.PersistentFlags().IntVar(&randomDelay, "limit-random-delay", 0, "Limit random delay")

	cmd.PersistentFlags().BoolVarP(&wordpress, "wordpress", "w", false, "Enable WordPress support")
	cmd.PersistentFlags().BoolVarP(&crypto, "crypto", "b", false, "Enable crypto features")
	cmd.PersistentFlags().BoolVarP(&email, "email", "e", false, "Enable email search")
	cmd.PersistentFlags().StringSliceVarP(&phone, "phone", "p", []string{}, "List of countries. RU,NL,DE,US. You can specify multiple or just one.")

	cli = configs.NewCLIConfig()
	cmd.PersistentFlags().VisitAll(cli.ConfigureViper("crawler"))
	CrawlCMD.AddCommand(SearchCMD)

}

func setCrawlerOptions(c *configs.Crawler, cmd *cobra.Command) {
	if cmd.Flags().Changed("tor") {
		c.Tor = tor
	}
	if cmd.Flags().Changed("url-revisit") {
		c.AllowURLRevisit = allowURLRevisit
	}
	if cmd.Flags().Changed("ignore-robots-txt") {
		c.IgnoreRobotsTxt = ignoreRobotsTxt
	}
	if cmd.Flags().Changed("debug") {
		c.Debug = debug
	}
	if cmd.Flags().Changed("wordpress") {
		c.Wordpress = wordpress
	}
	if cmd.Flags().Changed("email") {
		c.Email = email
	}
	if cmd.Flags().Changed("crypto") {
		c.Crypto = crypto
	}

	if cmd.Flags().Changed("urls") {
		c.StartingURLS = urls
	}
	if cmd.Flags().Changed("max-depth") {
		c.MaxDepth = maxDepth
	}
	if cmd.Flags().Changed("body-size") {
		c.MaxBodySize = maxBodySize
	}
	if cmd.Flags().Changed("cache-dir") {
		c.CacheDir = cacheDir
	}

	if cmd.Flags().Changed("queue-threads") {
		c.QueueThreads = queueThreads
	}

	if cmd.Flags().Changed("queue-max-size") {
		c.QueueMaxSize = queueMaxSize
	}

	if cmd.Flags().Changed("allowed-domains") {
		c.AllowedDomains = allowedDomains
	}

	if cmd.Flags().Changed("disallowed-url-filters") {
		c.DisallowedURLFilters = disallowedURLFilters
	}
	if cmd.Flags().Changed("disallowed-domains") {
		c.DisallowedDomains = disallowedDomains
	}
	if cmd.Flags().Changed("url-filters") {
		c.URLFilters = urlFilters
	}

	if cmd.Flags().Changed("limit-delay") {
		c.Delay = delay
	}

	if cmd.Flags().Changed("limit-random-delay") {
		c.RandomDelay = randomDelay
	}
	if cmd.Flags().Changed("phone") {
		c.PhoneNumbers = phone
	}

	logger.Infof("Wordpress: %t, Crypto: %t, Email: %t, Phone: %s", c.Wordpress, c.Crypto, c.Email, c.PhoneNumbers)

}
