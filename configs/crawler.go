package configs

type CollyConfig struct {
	StartingURLS         []string  `json:"startingUrls"`
	UserAgent            string    `json:"userAgent"`
	MaxDepth             int       `json:"maxDepth"`
	AllowedDomains       []string  `json:"allowedDomains"`
	DisallowedDomains    []string  `json:"disallowedDomains"`
	DisallowedURLFilters []string  `json:"disallowedURLFilters"`
	URLFilters           []string  `json:"urlFilters"`
	AllowURLRevisit      bool      `json:"allowURLRevisit"`
	MaxBodySize          int       `json:"maxBodySize"`
	CacheDir             string    `json:"cacheDir"`
	IgnoreRobotsTxt      bool      `json:"ignoreRobotsTxt"`
	QueueThreads         int       `json:"queueThreads"`
	QueueMaxSize         int       `json:"queueMaxSize"`
	Debug                bool      `json:"debug"`
	UseLimit             bool      `json:"useLimit"`
	LimitRule            LimitRule `json:"limitRule"`
}

type LimitRule struct {
	DomainRegexp string `json:"DomainRegexp"`
	Delay        int    `json:"Delay"`
	RandomDelay  int    `json:"RandomDelay"`
}

// Default is configs/json/crawlerConfig.json, chamge it to somewhere else if you need.
func loadCrawlerConfig() {
	var config CollyConfig

	loadConfig("configs/json/crawlerConfig.json", &config)
	cfg.CrawlerConf = config
}
