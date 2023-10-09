package crawler

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
	"github.com/r00tk3y/prying-deep/configs"
	"github.com/r00tk3y/prying-deep/pkg/logger"
	"github.com/r00tk3y/prying-deep/pkg/utils"
	"go.uber.org/zap"
)

func proxySetup(c *colly.Collector, tor configs.TorConfig) *colly.Collector {
	torProxy := fmt.Sprintf("socks5://%s:%s", tor.Host, tor.Port)
	rp, err := proxy.RoundRobinProxySwitcher(torProxy)
	if err != nil {
		logger.Fatal(err.Error())
	}
	c.SetProxyFunc(rp)

	checkTorConnection, err := utils.CheckIfTorConnectionExists(torProxy)
	if !checkTorConnection.IsTor {
		logger.Fatal("Killing session. Can't connect to Tor.\nError:  ", zap.Error(err))
	}
	return c

}
func NewCollector(config configs.CollyConfig, torConfig configs.TorConfig) *colly.Collector {
	maxBodySize := 1024 * 10 * 1024
	configBodySize := config.MaxBodySize * 10 * 1024
	c := colly.NewCollector()

	if config.UserAgent != "" {
		c.UserAgent = config.UserAgent
	}

	if config.MaxDepth != 0 {
		c.MaxDepth = config.MaxDepth
	}

	if len(config.AllowedDomains) != 0 {
		c.AllowedDomains = config.AllowedDomains
	}

	if len(config.DisallowedDomains) != 0 {
		c.DisallowedDomains = config.DisallowedDomains
	}

	if len(config.DisallowedURLFilters) != 0 {
		patterns, _ := utils.CompileRegexSlice(config.DisallowedURLFilters)
		c.DisallowedURLFilters = patterns
	}

	if len(config.URLFilters) != 0 {
		c.URLFilters = ConvertURLFiltersToRegexp(config.URLFilters)
	}

	if config.AllowURLRevisit {
		c.AllowURLRevisit = config.AllowURLRevisit
	}

	if maxBodySize != configBodySize {
		c.MaxBodySize = configBodySize
	}

	if config.CacheDir != "" {
		c.CacheDir = config.CacheDir
	}

	if config.IgnoreRobotsTxt {
		c.IgnoreRobotsTxt = config.IgnoreRobotsTxt
	}

	if config.Async {
		c.Async = config.Async
	}
	c = proxySetup(c, torConfig)
	return c
}
