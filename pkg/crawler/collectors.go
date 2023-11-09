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

// proxySetup initializes a new tor connection for the crawler to be able to parse onion links
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

// NewCollector initializes colly.NewCollector with the modifications needed for extravagant crawling.
func NewCollector(config configs.Crawler, torConfig configs.TorConfig) *colly.Collector {
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
		patterns, err := utils.CompileRegex(config.DisallowedURLFilters)
		if err != nil {
			color.HiRed("Error.. Please check your regexp in DissalowedURLFilters")
			log.Fatal(err)
		}
		c.DisallowedURLFilters = patterns
	}

	if len(config.URLFilters) != 0 {
		filters, err := utils.CompileRegex(config.URLFilters)
		if err != nil {
			color.HiRed("Error.. Please check your regexp in URLFilters")
			log.Fatal(err)
		}
		c.URLFilters = filters

	}

	if config.AllowURLRevisit {
		c.AllowURLRevisit = config.AllowURLRevisit
	}

	if maxBodySize != configBodySize {
		c.MaxBodySize = configBodySize
	}

	//I'm not sure if this works for now
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
	c.DetectCharset = true
	//TODO: add this option into the config
	c.SetRequestTimeout(time.Second * 30)
	c = proxySetup(c, torConfig)
	return c
}
