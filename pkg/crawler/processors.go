package crawler

import (
	"github.com/r00tk3y/prying-deep/models"
	"github.com/r00tk3y/prying-deep/pkg/logger"
	"github.com/r00tk3y/prying-deep/pkg/pryingtools/email"
	"github.com/r00tk3y/prying-deep/pkg/pryingtools/phonenumber"
	"github.com/r00tk3y/prying-deep/pkg/pryingtools/wordpress"
)

func processWordPress(body, url string, pageId int) {
	wordpressMatches, _ := wordpress.FindWordpressPatterns(body, url)
	if len(wordpressMatches) != 0 {
		logger.Infof("Number of WordPress matches: %d", len(wordpressMatches))
		models.CreateWordPressFootPrint(pageId, wordpressMatches)
	}
}

func processEmail(body string, pageId int) {
	emailMatches := email.FindEmail(body)
	if len(emailMatches) != 0 {
		logger.Infof("Email matches: %s", emailMatches)
		models.CreateEmails(pageId, emailMatches)
	}
}

func processPhones(body string, pageId int, countryMaps map[string]string) {
	phoneProcessor := phonenumber.NewPhoneProcessor()
	phoneProcessor.ProcessPhoneNumbers(body, pageId, countryMaps)
}
