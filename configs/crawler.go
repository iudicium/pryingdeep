package configs

// Detailed documentation for different options is here.
// https://github.com/gocolly/colly/blob/3c987f1982edbb5ba8876eef56dd35e1ff05932a/colly.go#L55C24-L55C24

type CollyConfig struct {
	// StartingURLS is the entry point urls
	StartingURLS []string `json:"startingUrls"`
	// UserAgent is the User-Agent string used by HTTP requests
	UserAgent string `json:"userAgent"`
	// MaxDepth limits the recursion depth of visited URLs.
	// Set it to 0 for infinite recursion (default).
	MaxDepth int `json:"maxDepth"`
	//AllowedDomains is a domain whitelist
	AllowedDomains []string `json:"allowedDomains"`
	// DisallowedDomains  is a domain blacklist
	DisallowedDomains []string `json:"disallowedDomains"`
	// DisallowedURLFilters is a list of regular expressions which restricts
	// visiting URLs. If any of the rules matches to a URL the
	// request will be stopped. DisallowedURLFilters will
	// be evaluated before URLFilters
	// Leave it blank to allow any URLs to be visited
	DisallowedURLFilters []string `json:"disallowedURLFilters"`
	// URLFilters is a list of regular expressions which restricts
	// visiting URLs. If any of the rules matches to a URL the
	// request won't be stopped. DisallowedURLFilters will
	// be evaluated before URLFilters

	// Leave it blank to allow any URLs to be visited
	URLFilters []string `json:"urlFilters"`
	// AllowURLRevisit allows multiple downloads of the same URL

	AllowURLRevisit bool `json:"allowURLRevisit"`
	// MaxBodySize is the limit of the retrieved response body in bytes.
	// 0 means unlimited.
	// The default value for MaxBodySize is 10MB (10 * 1024 * 1024 bytes).

	MaxBodySize int `json:"maxBodySize"`
	// CacheDir specifies a location where GET requests are cached as files.
	// When it's not defined, caching is disabled.

	CacheDir string `json:"cacheDir"`
	// IgnoreRobotsTxt allows the Collector to ignore any restrictions set by
	// the target host's robots.txt file.  See http://www.robotstxt.org/ for more
	// information.

	IgnoreRobotsTxt bool `json:"ignoreRobotsTxt"`

	//QueueThreads defines the number of consumer threads
	QueueThreads int `json:"queueThreads"`
	//QueueMaxSize  defines the capacity of the queue.
	// New requests are discarded if the queue size reaches MaxSize
	QueueMaxSize int `json:"queueMaxSize"`

	Debug bool `json:"debug"`
	//UseLimit tells the crawler wherether to use LimitRule or not
	UseLimit  bool      `json:"useLimit"`
	LimitRule LimitRule `json:"limitRule"`
}

type LimitRule struct {
	DomainRegexp string `json:"DomainRegexp"`
	Delay        int    `json:"Delay"`
	RandomDelay  int    `json:"RandomDelay"`
}

// loadCrawlerConfig is used to load the configs from json and apply them to the Config struct
func loadCrawlerConfig() {
	var config CollyConfig
	//Probably best to not hard code the values
	loadConfig("configs/json/crawlerConfig.json", &config)
	cfg.CrawlerConf = config
}
