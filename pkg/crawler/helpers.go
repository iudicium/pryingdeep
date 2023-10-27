package crawler

import (
	"regexp"

	"github.com/gocolly/colly/v2"
	"github.com/pryingbytez/prying-deep/configs"
	"github.com/pryingbytez/prying-deep/models"
	"github.com/pryingbytez/prying-deep/pkg/logger"
	"github.com/pryingbytez/prying-deep/pkg/utils"
)

func HandleResponse(response *colly.Response, options *configs.PryingConfig) {
	body := string(response.Body)
	url := response.Request.URL.String()
	logger.Infof("Crawling url: %s", url)
	pageId, err := ParseResponse(url, body, response)

	if err != nil {
		logger.Errorf("Something went wrong during parsing the response from: %s\nErr: %s ", url, err)
	}

	if options.Wordpress {
		go processWordPress(body, pageId)
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

func ParseResponse(url string, body string, response *colly.Response) (int, error) {
	title, _ := utils.ExtractTitleFromBody(body)
	headers := utils.CreateMapFromValues(*response.Headers)

	ResId, err := models.CreatePage(
		url,
		title,
		response.StatusCode,
		body,
		headers,
	)
	if err != nil {
		return 0, err
	}

	return int(ResId), nil
}
