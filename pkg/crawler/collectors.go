package crawler

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
	"github.com/gocolly/colly/v2/proxy"
	"go.uber.org/zap"

	"github.com/pryingbytez/pryingdeep/configs"
	"github.com/pryingbytez/pryingdeep/pkg/logger"
	"github.com/pryingbytez/pryingdeep/pkg/utils"
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

// TODO: add a command line interface/UI on web to process this
func NewCollector(config configs.CollyConfig, torConfig configs.TorConfig) *colly.Collector {
	if len(config.StartingURLS) == 0 {
		color.Red("Exiting.. No starting urls were provided")
		os.Exit(0)
	}
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

	if config.Debug {
		c.SetDebugger(&debug.LogDebugger{})
	}
	if config.UseLimit {
		//TODO: add support for all settings maybe
		var rule colly.LimitRule
		if config.LimitRule.Delay != 0 {
			rule.Delay = time.Duration(config.LimitRule.Delay) * time.Second
		}

		if config.LimitRule.RandomDelay != 0 {
			rule.RandomDelay = time.Duration(config.LimitRule.RandomDelay) * time.Second
		}
		if config.LimitRule.DomainRegexp != "" {
			rule.DomainRegexp = config.LimitRule.DomainRegexp

			err := c.Limit(&rule)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	//TODO: add this option into the config
	c.SetRequestTimeout(time.Second * 30)
	c = proxySetup(c, torConfig)
	return c
}
