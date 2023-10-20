package tests

import (
	"fmt"
	"github.com/r00tk3y/prying-deep/configs"
	"github.com/r00tk3y/prying-deep/models"
	"github.com/r00tk3y/prying-deep/pkg/logger"
	"github.com/r00tk3y/prying-deep/pkg/pryingtools/phonenumber"
	"github.com/r00tk3y/prying-deep/pkg/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type PhoneNumberValidationTestConfig struct {
	Filename      string
	Regex         string
	CountryCode   string
	ExpectedCount int
	WebPageId     int
}

func TestSetup(t *testing.T) {
	configs.SetupEnvironment()
	cfg := configs.GetConfig().DbConf

	logger.InitLogger()
	defer logger.Logger.Sync()

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbTestName)
	models.SetupDatabase(dbURL)
}

func testPhoneNumberValidation(t *testing.T, testConfig PhoneNumberValidationTestConfig) {
	assert := assert.New(t)
	html := utils.ReadFile(testConfig.Filename)

	testPatterns := map[string]string{testConfig.CountryCode: testConfig.Regex}

	phoneProcessor := phonenumber.NewPhoneProcessor()
	phoneProcessor.SetCountryRegex(testConfig.CountryCode, testConfig.Regex)

	phoneProcessor.ProcessPhoneNumbers(html, testConfig.WebPageId, testPatterns)

	phones, err := models.GetPhoneNumbers(testConfig.WebPageId)
	if err != nil {
		logger.Errorf("err getting phones from db: %s", err)
	}

	assert.Equal(len(phones), testConfig.ExpectedCount)
	t.Cleanup(func() {
		models.DeletePhoneNumbersByCountryCode(testConfig.CountryCode)
		t.Log(fmt.Sprintf("Performing %s numbers clean up...", testConfig.CountryCode))
	})
}
func TestRussianPhoneNumberValidation(t *testing.T) {
	testConfig := PhoneNumberValidationTestConfig{
		Filename:      "ru.html",
		Regex:         phonenumber.RuRegex,
		CountryCode:   "RU",
		ExpectedCount: 22,
		WebPageId:     1,
	}
	testPhoneNumberValidation(t, testConfig)
}

func TestUSAPhoneNumberValidation(t *testing.T) {
	testConfig := PhoneNumberValidationTestConfig{
		Filename:      "usa.html",
		Regex:         phonenumber.USRegex,
		CountryCode:   "US",
		ExpectedCount: 151,
		WebPageId:     1,
	}
	testPhoneNumberValidation(t, testConfig)
}

//
//There's 15-16 duplicate numbers in the html, so we store only 15

func TestUKPhoneNumberValidation(t *testing.T) {
	testConfig := PhoneNumberValidationTestConfig{
		Filename:      "uk.html",
		Regex:         phonenumber.UKRegex,
		CountryCode:   "GB",
		ExpectedCount: 15,
		WebPageId:     1,
	}
	testPhoneNumberValidation(t, testConfig)
}

// // DE = Germany
// // 5 Duplicates, only 1 Correct
func TestDEPhoneNumberValidation(t *testing.T) {
	testConfig := PhoneNumberValidationTestConfig{
		Filename:      "de.html",
		Regex:         phonenumber.DERegex,
		CountryCode:   "DE",
		ExpectedCount: 1,
		WebPageId:     1,
	}
	testPhoneNumberValidation(t, testConfig)
}

// FIXME: the nl regexp validation is a bit wrong but it works for now

func TestNLPhoneNumberValidation(t *testing.T) {
	testConfig := PhoneNumberValidationTestConfig{
		Filename:      "nl.html",
		Regex:         phonenumber.NLRegex,
		CountryCode:   "NL",
		ExpectedCount: 2,
		WebPageId:     1,
	}
	testPhoneNumberValidation(t, testConfig)
}
func TestConcurrentPhoneProcessing(t *testing.T) {
	html, err := utils.ReadFilesInDirectory("data")
	if err != nil {
		t.Fatal(err)
	}
	phoneProcessor := phonenumber.NewPhoneProcessor()

	testCountryRegexPatterns := map[string]string{
		"RU": phonenumber.RuRegex,
		"US": phonenumber.USRegex,
		"GB": phonenumber.UKRegex,
		"DE": phonenumber.DERegex,
		"NL": phonenumber.NLRegex,
	}

	phoneProcessor.ProcessPhoneNumbers(html, 1, testCountryRegexPatterns)
	t.Cleanup(func() {
		for countryCode, _ := range testCountryRegexPatterns {
			logger.Debugf("removing %s phones from test database", countryCode)
			models.DeletePhoneNumbersByCountryCode(countryCode)

		}
	})

}
