package configs

import (
	"encoding/json"
	"log"
	"os"
)

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

func LoadCrawlerConfig() {
	log.Println("Loading crawler config...")
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
