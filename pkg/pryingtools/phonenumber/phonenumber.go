package phonenumber

import (
	"fmt"
	"github.com/r00tk3y/prying-deep/models"
	"github.com/r00tk3y/prying-deep/pkg/logger"
	"regexp"

	"github.com/nyaruka/phonenumbers"
)

const (
	RuRegex = `(^8|7|\+7)((\d{10})|(\s\(\d{3}\)\s\d{3}\s\d{2}\s\d{2}))`
	USRegex = `\(?\d{3}\)?-? *\d{3}-? *-?\d{4}`
)

// PhoneNumberValidator is an implementation of PhoneNumberMatcher for a specific country.
type PhoneNumberFinder interface {
	FindPhoneNumbers(html string) []string
	FormatAndCreateNumbers(webPageId int, phoneNumbers []string) error
}

type PhoneNumberValidator struct {
	regex       *regexp.Regexp
	countryCode string
}

// NewPhoneNumberValidator creates a new instance of PhoneNumberValidator.
func NewPhoneNumberValidator(regexPattern string, countryCode string) (*PhoneNumberValidator, error) {
	regex, err := regexp.Compile(regexPattern)
	if err != nil {
		return nil, err
	}
	return &PhoneNumberValidator{regex, countryCode}, nil
}
func (c *PhoneNumberValidator) FindPhoneNumbers(html string) []string {
	return c.regex.FindAllString(html, -1)
}

// Validates  phone numbers and returns a list of validated phone numbers with as much information as possible
// From libphonenumbers
// Function is not perfect and may sometimes fail
func (p *PhoneNumberValidator) FormatAndCreateNumbers(webPageId int, phoneNumbers []string) error {
	for _, phoneNumber := range phoneNumbers {
		fmt.Println(phoneNumber)
		num, err := phonenumbers.Parse(phoneNumber, p.countryCode)

		if err != nil {
			logger.Errorf("err during parsing phone number: %s", err)
		}

		if phonenumbers.IsValidNumber(num) {

			interNum := phonenumbers.Format(num, phonenumbers.INTERNATIONAL)
			NatNum := phonenumbers.Format(num, phonenumbers.NATIONAL)
			fmt.Println(interNum, NatNum)
			err = models.CreatePhoneNumber(webPageId, interNum, NatNum, p.countryCode)
			if err != nil {
				//logger.Errorf("error during creation of phone numbers; %s", err)
			}
		}
	}
	return nil
}

// TODO: it looks like i will have to create regex patterns for each country.
// TODO: Let's add support for 5 countries.
// TODO: Russia, USA, UK, Germany, Netherlands
