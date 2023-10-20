package crawler

import (
	"github.com/gocolly/colly/v2"
	"github.com/r00tk3y/prying-deep/configs"
	"github.com/r00tk3y/prying-deep/pkg/logger"
	"github.com/r00tk3y/prying-deep/pkg/parsers"
	"regexp"
)

func HandleResponse(response *colly.Response, options *configs.PryingConfig) {
	body := string(response.Body)
	url := response.Request.URL.String()
	logger.Infof("Crawling url: %s", url)
	pageId, err := parsers.ParseResponse(url, body, response)

	if err != nil {
		logger.Errorf("Something went wrong during parsing the response from: %s\nErr: %s ", url, err)
	}

	if options.Wordpress {
		go processWordPress(body, url, pageId)
	}

	if options.Email {
		go processEmail(body, pageId)
	}
	if options.Crypto {
		go processCrypto(body, pageId)
	}
	if len(options.PhoneNumbers) != 0 {
		go processPhones(body, pageId, options.PhoneNumbers)
	}

}
func ConvertURLFiltersToRegexp(filters []string) []*regexp.Regexp {
	var urlFilters []*regexp.Regexp

	for _, filter := range filters {
		regex := regexp.MustCompile(filter)
		urlFilters = append(urlFilters, regex)
	}

	return urlFilters
}
