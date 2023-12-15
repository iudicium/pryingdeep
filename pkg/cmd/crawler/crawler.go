package crawler

import (
	"errors"
	"fmt"
	"os"

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
		//Empty list, we don't want to override the actual urls that the user passes
		Crawl([]string{})
	},
}

func setupCrawlerConfig(cmd *cobra.Command) *configs.Configuration {
	configs.SetupEnvironment()
	cfg = configs.GetConfig()
	setCrawlerOptions(&cfg.Crawler, cmd)
	return cfg
}

func Crawl(urls []string) {
	newCrawler := crawler.NewCrawler(cfg.Tor, cfg.Crawler)
	if err := newCrawler.Crawl(); err != nil {
		handleCrawlError(err)
	}
}

func init() {
	initCrawler(CrawlCMD, "default")
}

func initCrawler(cmd *cobra.Command, crawlerType string) {
	cmd.Flags().StringSliceVarP(&urls, "urls", "u", nil, "Entry point URLs")
	cmd.Flags().BoolVarP(&tor, "tor", "t", false, "Turn on/off connecting to Tor.")
	cmd.Flags().StringVar(&userAgent, "user-agent", userAgent, "Specify any user agents for the crawler to use.")
	cmd.Flags().IntVar(&maxDepth, "max-depth", 0, "Maximum recursion depth")
	cmd.Flags().IntVar(&maxBodySize, "body-size", 0, "Max body size in bytes (0 for unlimited)")
	cmd.Flags().StringVar(&cacheDir, "cache-dir", "", "Cache directory")
	cmd.Flags().BoolVar(&ignoreRobotsTxt, "ignore-robots-txt", false, "Ignore robots.txt")
	cmd.Flags().BoolVar(&debug, "debug", false, "Enable debug mode")
	cmd.Flags().IntVar(&queueThreads, "queue-threads", 4, "Number of queue threads")
	cmd.Flags().IntVar(&queueMaxSize, "queue-max-size", 50000, "Queue max size")

	cmd.Flags().StringSliceVar(&allowedDomains, "allowed-domains", nil, "Allowed domains")
	cmd.Flags().StringSliceVar(&disallowedDomains, "disallowed-domains", nil, "Disallowed domains")
	cmd.Flags().StringSliceVar(&disallowedURLFilters, "disallowed-url-filters", nil, "Disallowed URL filters")
	cmd.Flags().StringSliceVar(&urlFilters, "url-filters", nil, "URL filters")
	cmd.Flags().BoolVar(&allowURLRevisit, "url-revisit", false, "Allow URL revisit")

	cmd.Flags().IntVar(&delay, "limit-delay", 0, "Limit delay")
	cmd.Flags().IntVar(&randomDelay, "limit-random-delay", 0, "Limit random delay")

	cmd.Flags().BoolVarP(&wordpress, "wordpress", "w", false, "Enable WordPress support")
	cmd.Flags().BoolVarP(&crypto, "crypto", "b", false, "Enable crypto features")
	cmd.Flags().BoolVarP(&email, "email", "e", false, "Enable email search")
	cmd.Flags().StringSliceVarP(&phone, "phone", "p", []string{}, "List of countries. RU,NL,DE,US. You can specify multiple or just one.")

	switch crawlerType {
	case "search":
		{
			cmd.Flags().StringSliceVarP(&keywords, "keywords", "k", nil, "List of keywords or sentences for search")
			cmd.Flags().MarkHidden("urls")

		}
	}

	cli := configs.NewCLIConfig()
	cmd.Flags().VisitAll(cli.ConfigureViper("crawler"))

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

	if cmd.Flags().Changed("keywords") {
		c.Keywords = keywords
		generateSearchURLS(keywords)
	} else {
		if len(c.Keywords) == 0 {
			fmt.Println(color.RedString("No keywords were provided while using the search command."))
			os.Exit(1)

		}
	}

	logger.Infof("Wordpress: %t, Crypto: %t, Email: %t, Phone: %s", c.Wordpress, c.Crypto, c.Email, c.PhoneNumbers)

}

func handleCrawlError(err error) {
	if errors.Is(err, colly.ErrQueueFull) {
		color.HiRed("\nQueue max size has been reached! Exiting.")
	} else {
		logger.Errorf("Crawl error: %s", err)
	}
}
