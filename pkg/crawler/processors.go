package crawler

import (
	"github.com/iudicium/pryingdeep/models"
	"github.com/iudicium/pryingdeep/pkg/logger"
	"github.com/iudicium/pryingdeep/pkg/pryingtools/cryptoscanner"
	"github.com/iudicium/pryingdeep/pkg/pryingtools/email"
	"github.com/iudicium/pryingdeep/pkg/pryingtools/phonenumber"
	"github.com/iudicium/pryingdeep/pkg/pryingtools/wordpress"
)

//This file is used primarily for our prying modules, so that we can create new goroutines based on the specified options!

func processWordPress(body string, pageId int) {
	wordpressMatches, _ := wordpress.FindWordpressPatterns(body)
	if len(wordpressMatches) != 0 {
		logger.Infof("Number of WordPress matches: %d", len(wordpressMatches))
		models.CreateWordPressFootPrint(pageId, wordpressMatches)
	}
}

func processEmail(body string, pageId int) {
	emailFinder := email.NewEmailFinder()

	emailMatches := emailFinder.FindEmails(body)
	if len(emailMatches) != 0 {
		logger.Infof("Email matches: %s", emailMatches)
		models.CreateEmails(pageId, emailMatches)
	}
}

func processPhones(body string, pageId int, countryCodes []string) {
	countryMaps := phonenumber.MapCountryCodeToRegex(countryCodes)
	phoneProcessor := phonenumber.NewPhoneProcessor()
	phoneProcessor.ProcessPhoneNumbers(body, pageId, countryMaps)
}

func processCrypto(body string, pageId int) {
	crypto := cryptoscanner.NewCryptoScanner()
	crypto.Search(body, pageId)
}
