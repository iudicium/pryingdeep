package configs

// Detailed documentation for different options is here.
// https://github.com/gocolly/colly/blob/3c987f1982edbb5ba8876eef56dd35e1ff05932a/colly.go#L55C24-L55C24

type Crawler struct {
	// StartingURLS is the entry point urls
	StartingURLS []string `mapstructure:"urls"`
	Tor          bool     `mapstructure:"tor"`
	// UserAgent is the User-Agent string used by HTTP requests
	UserAgent string `mapstructure:"user-agent"`
	// MaxDepth limits the recursion depth of visited URLs.
	// Set it to 0 for infinite recursion (default).
	MaxDepth int `mapstructure:"max-depth"`

	// AllowedDomains is a domain whitelist
	AllowedDomains []string `mapstructure:"allowed-domains"`

	// DisallowedDomains is a domain blacklist
	DisallowedDomains []string `mapstructure:"disallowed-domains"`

	// DisallowedURLFilters is a list of regular expressions which restricts
	// visiting URLs. If any of the rules matches to a URL the
	// request will be stopped. DisallowedURLFilters will
	// be evaluated before URLFilters
	// Leave it blank to allow any URLs to be visited
	DisallowedURLFilters []string `mapstructure:"disallowed-url-filters"`

	// URLFilters is a list of regular expressions which restricts
	// visiting URLs. If any of the rules matches to a URL the
	// request won't be stopped. DisallowedURLFilters will
	// be evaluated before URLFilters
	// Leave it blank to allow any URLs to be visited
	URLFilters []string `mapstructure:"url-filters"`

	// AllowURLRevisit allows multiple downloads of the same URL
	AllowURLRevisit bool `mapstructure:"allow-url-revisit"`
	// MaxBodySize is the limit of the retrieved response body in bytes.
	// 0 means unlimited.
	// The default value for MaxBodySize is 10MB (10 * 1024 * 1024 bytes)
	MaxBodySize int `mapstructure:"body-size"`
	// CacheDir specifies a location where GET requests are cached as files.
	// When it's not defined, caching is disabled.
	CacheDir string `mapstructure:"cache-dir"`
	// IgnoreRobotsTxt allows the Collector to ignore any restrictions set by
	// the target host's robots.txt file. See http://www.robotstxt.org/ for more
	// information.
	IgnoreRobotsTxt bool `mapstructure:"ignore-robots-txt"`
	// QueueThreads defines the number of consumer threads
	QueueThreads int `mapstructure:"queue-threads"`
	// QueueMaxSize defines the capacity of the queue.
	// New requests are discarded if the queue size reaches MaxSize
	QueueMaxSize int  `mapstructure:"queue-max-size"`
	Debug        bool `mapstructure:"debug"`

	Delay       int `mapstructure:"delay"`
	RandomDelay int `mapstructure:"random-delay"`

	//PhoneNumbers List of countries. RU,NL,DE,GB,US. You can specify multiple or just one.
	//Default is blank
	PhoneNumbers []string `mapstructure:"phone-numbers"`
	Email        bool     `mapstructure:"email"`
	Crypto       bool     `mapstructure:"crypto"`
	Wordpress    bool     `mapstructure:"wordpress"`
}
type LimitRule struct {
}

func loadCrawlerConfig() {
	var config Crawler
	loadConfig("crawler", &config)
	cfg.Crawler = config
}
