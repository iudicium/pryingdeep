package crawler

import (
	"github.com/iudicium/pryingdeep/configs"
)

var (
	cfg *configs.Configuration
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
