package crawler

import (
	"github.com/gocolly/colly/v2"
	"github.com/r00tk3y/prying-deep/models"
	"github.com/r00tk3y/prying-deep/pkg/logger"
	"github.com/r00tk3y/prying-deep/pkg/parsers"
	"github.com/r00tk3y/prying-deep/pkg/pryingtools/email"
	"github.com/r00tk3y/prying-deep/pkg/pryingtools/phonenumber"
	"github.com/r00tk3y/prying-deep/pkg/pryingtools/wordpress"
	"regexp"
)

func HandleResponse(response *colly.Response) {
	body := string(response.Body)

	url := response.Request.URL.String()
	logger.Infof("Crawling url: %s", url)
	pageId, err := parsers.ParseResponse(url, body, response)
	if err != nil {
		logger.Errorf("Something went wrong during parsing the response from: %s\nErr: %s ", url, err)
	}
	wordpressMatches, _ := wordpress.FindWordpressPatterns(body, url)
	if len(wordpressMatches) != 0 {
		logger.Infof("Number of Wordpress matches: %d", len(wordpressMatches))
		models.CreateWordPressFootPrint(int(pageId), wordpressMatches)
	}

	emailMatches := email.FindEmail(body)
	if len(emailMatches) != 0 {
		logger.Infof("Email matches: %s", emailMatches)
		models.CreateEmails(int(pageId), emailMatches)
	}
	//For now this works, later on i will have to add commands to run through all of em
	phoneValidator, err := phonenumber.NewPhoneNumberValidator(phonenumber.RuRegex, "RU")
	phones := phoneValidator.FindPhoneNumbers(body)
	err = phoneValidator.FormatAndCreateNumbers(int(pageId), phones)
	if err != nil {
		logger.Errorf("something has gone wrong: %s", err)
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
