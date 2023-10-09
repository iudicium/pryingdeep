package configs

import (
	"encoding/json"
	"log"
	"os"
)

type CollyConfig struct {
	StartingURL          string   `json:"startingUrl"`
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
	Async                bool     `json:"async"`
	QueueThreads         int      `json:"queueThreads"`
	QueueMaxSize         int      `json:"queueMaxSize"`
}

func LoadCrawlerConfig() {
	var config CollyConfig

	fileContent, err := os.ReadFile("crawlerConfig.json")
	if err != nil {
		log.Println("error during loading crawler config:", err)
	}

	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		log.Println("error during loading crawler config:", err)
	}
	cfg.CrawlerConf = config
}
