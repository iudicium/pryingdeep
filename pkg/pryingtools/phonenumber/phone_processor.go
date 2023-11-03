package phonenumber

import (
	"sync"

	"github.com/pryingbytez/pryingdeep/pkg/logger"
)

type PhoneProcessor struct {
	regexPatterns map[string]string
}

func NewPhoneProcessor() *PhoneProcessor {
	return &PhoneProcessor{
		regexPatterns: make(map[string]string),
	}
}

func (p *PhoneProcessor) SetCountryRegex(countryCode, regexPattern string) {
	p.regexPatterns[countryCode] = regexPattern
}

func (p *PhoneProcessor) ProcessPhoneNumbers(html string, webPageID int, patterns map[string]string) {
	var wg sync.WaitGroup

	for countryCode, regexPattern := range patterns {
		wg.Add(1)
		go func(countryCode, regexPattern string) {
			defer wg.Done()

			validator, err := NewPhoneNumberValidator(regexPattern, countryCode)
			if err != nil {
				logger.Errorf("establishing NewPhoneValidator has gone wrong: %s", err)
			}
			phones := validator.FindPhoneNumbers(html)
			if len(phones) != 0 {
				logger.Infof("Found %d phone numbers for counryCode: %s", len(phones), countryCode)
				validator.FormatAndCreateNumbers(webPageID, phones)
			}
		}(countryCode, regexPattern)
	}

	wg.Wait()
}
