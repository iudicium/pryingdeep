package phonenumber

import (
	"regexp"
	"strings"

	"github.com/iudicium/pryingdeep/models"
	"github.com/iudicium/pryingdeep/pkg/logger"

	"github.com/nyaruka/phonenumbers"
)

const (
	RuRegex = `(^8|7|\+7)((\d{10})|(\s\(\d{3}\)\s\d{3}\s\d{2}\s\d{2}))`
	USRegex = `\(?\d{3}\)?-? *\d{3}-? *-?\d{4}`
	DERegex = `[^\d]((\+49|0049|0)[\s]?1[567]\d{1,2}([ \-/]*\d){7})`
	NLRegex = `(?:(?:\+31|0|0031)[\s-]?\d{1,3}[\s-]?\d{6,7}|06[\s-]?\d{8})`
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

// FormatAndCreateNumbers Validates phone numbers
// and returns a list of validated phone numbers with as much information as possible
// From libphonenumbers
// Function is not perfect and may sometimes fail
func (p *PhoneNumberValidator) FormatAndCreateNumbers(webPageId int, phoneNumbers []string) {
	for _, phoneNumber := range phoneNumbers {
		num, err := phonenumbers.Parse(phoneNumber, p.countryCode)
		if err != nil {
			logger.Errorf("err during parsing phone number: %sCountry Code: %s", err, p.countryCode)
		}
		if phonenumbers.IsValidNumber(num) {
			logger.Infof("Valid num: %s", num.String())
			interNum := phonenumbers.Format(num, phonenumbers.INTERNATIONAL)
			NatNum := phonenumbers.Format(num, phonenumbers.NATIONAL)
			err = models.CreatePhoneNumber(webPageId, interNum, NatNum, p.countryCode)
			if err != nil {
				logger.Errorf("error during creation of phone numbers; %s", err)

			}
		}
	}

}

// FindTelTagPhones is the primary function for searching for phone numbers in a webpage.
// It can search for href tel tags and input type tel tags, combining them and returning results
// separated by a space as a string for further classification of the phone number.
func FindTelTagPhones(html string) (string, error) {
	regex := `<(?:a|input)[^>]*(?:\s+href\s*=\s*["']\s*tel:([^"']*)["']|type\s*=\s*["']\s*tel["'][^>]*)>`
	validator, err := NewPhoneNumberValidator(regex, "")
	if err != nil {
		return "", err
	}
	phones := validator.FindPhoneNumbers(html)
	phoneHTML := strings.Join(phones, " ")
	return phoneHTML, nil
}
