package tests

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pryingbytez/pryingdeep/configs"
	"github.com/pryingbytez/pryingdeep/models"
	"github.com/pryingbytez/pryingdeep/pkg/logger"
	"github.com/pryingbytez/pryingdeep/pkg/pryingtools/phonenumber"
	"github.com/pryingbytez/pryingdeep/pkg/utils"
)

type PhoneNumberValidationTestConfig struct {
	URL           string
	Regex         string
	CountryCode   string
	ExpectedCount int
	WebPageId     int
}

func TestSetup(t *testing.T) {
	configs.SetupEnvironment()
	cfg := configs.GetConfig().DbConf

	logger.InitLogger(false)
	defer logger.Logger.Sync()

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbTestName)
	models.SetupDatabase(dbURL)
}

func testPhoneNumberValidation(t *testing.T, testConfig PhoneNumberValidationTestConfig) {
	assert := assert.New(t)

	resp, err := http.Get(testConfig.URL)
	if err != nil {
		t.Fatalf("Error making GET request: %v", err)
	}
	defer resp.Body.Close()

	// Read the HTML content from the response
	html, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}

	testPatterns := map[string]string{testConfig.CountryCode: testConfig.Regex}

	phoneProcessor := phonenumber.NewPhoneProcessor()
	phoneProcessor.SetCountryRegex(testConfig.CountryCode, testConfig.Regex)

	phoneProcessor.ProcessPhoneNumbers(string(html), testConfig.WebPageId, testPatterns)

	phones, err := models.GetPhoneNumbers(testConfig.WebPageId)
	if err != nil {
		logger.Errorf("Error getting phones from the database: %s", err)
	}

	assert.Equal(len(phones), testConfig.ExpectedCount)

	t.Cleanup(func() {
		models.DeletePhoneNumbersByCountryCode(testConfig.CountryCode)
		t.Log(fmt.Sprintf("Performing %s numbers clean up...", testConfig.CountryCode))
	})
}

func TestRussianPhoneNumberValidation(t *testing.T) {
	testConfig := PhoneNumberValidationTestConfig{
		URL:           "https://mysmsbox.ru/",
		Regex:         phonenumber.RuRegex,
		CountryCode:   "RU",
		ExpectedCount: 47,
		WebPageId:     1,
	}
	testPhoneNumberValidation(t, testConfig)
}
func TestUSAPhoneNumberValidation(t *testing.T) {
	testConfig := PhoneNumberValidationTestConfig{
		URL:           "https://www.thisnumber.com/270-258",
		Regex:         phonenumber.USRegex,
		CountryCode:   "US",
		ExpectedCount: 151,
		WebPageId:     1,
	}
	testPhoneNumberValidation(t, testConfig)
}

// //
// //There's 15-16 duplicate numbers in the html, so we store only 15
//
// // // DE = Germany
// // // 5 Duplicates, only 1 Correct
func TestDEPhoneNumberValidation(t *testing.T) {
	testConfig := PhoneNumberValidationTestConfig{
		URL:           "https://allaboutberlin.com/guides/dial-phone-numbers-germany",
		Regex:         phonenumber.DERegex,
		CountryCode:   "DE",
		ExpectedCount: 1,
		WebPageId:     1,
	}
	testPhoneNumberValidation(t, testConfig)
}

// // FIXME: the nl regexp validation is a bit wrong but it works for now
func TestConcurrentPhoneProcessing(t *testing.T) {
	html, err := utils.ReadFilesInDirectory("data")
	if err != nil {
		t.Fatal(err)
	}
	phoneProcessor := phonenumber.NewPhoneProcessor()

	testCountryRegexPatterns := map[string]string{
		"GB": phonenumber.UKRegex,
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
