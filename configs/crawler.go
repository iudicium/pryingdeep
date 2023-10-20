package configs

type CollyConfig struct {
	StartingURLS         []string `json:"startingUrls"`
	UserAgent            string   `json:"userAgent"`
	MaxDepth             int      `json:"maxDepth"`
	AllowedDomains       []string `json:"allowedDomains"`
	DisallowedDomains    []string `json:"disallowedDomains"`
	DisallowedURLFilters []string `json:"disallowedURLFilters"`
	URLFilters           []string `json:"urlFilters"`
	AllowURLRevisit      bool     `json:"allowURLRevisit"`
	MaxBodySize          int      `json:"maxBodySize"`
	CacheDir             string   `json:"cacheDir"`
	IgnoreRobotsTxt      bool     `json:"ignoreRobotsTxt"`
	QueueThreads         int      `json:"queueThreads"`
	QueueMaxSize         int      `json:"queueMaxSize"`
	Debug                bool     `json:"debug"`
}

func loadCrawlerConfig() {
	var config CollyConfig
	loadConfig("crawlerConfig.json", &config)
	cfg.CrawlerConf = config
}
