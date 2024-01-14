package phonenumber

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/iudicium/pryingdeep/models"
	"github.com/iudicium/pryingdeep/pkg/fsutils"
	"github.com/iudicium/pryingdeep/pkg/logger"
	"github.com/iudicium/pryingdeep/tests/test_helpers"
)

var db *gorm.DB

type PhoneNumberValidationTestConfig struct {
	URL           string
	Regex         string
	CountryCode   string
	ExpectedCount int
	WebPageId     int
}

func ReadFilesInDirectory(directoryPath string) (string, error) {
	files, err := os.ReadDir(directoryPath)
	if err != nil {
		return "", err
	}

	var result string

	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(directoryPath, file.Name())
			content, err := os.ReadFile(filePath)
			if err != nil {
				return "", err
			}
			result += string(content)
		}
	}

	return result, nil
}

func TestSetup(t *testing.T) {
	test_helpers.InitTestConfig()
	test_helpers.CreateTestWebPage()
	db = models.GetDB()
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

	phoneProcessor := NewPhoneProcessor()
	phoneProcessor.SetCountryRegex(testConfig.CountryCode, testConfig.Regex)

	phoneProcessor.ProcessPhoneNumbers(string(html), testConfig.WebPageId, testPatterns)

	phones, err := test_helpers.GetPhoneNumbers(db, testConfig.WebPageId)
	if err != nil {
		logger.Errorf("Error getting phones from the database: %s", err)
	}

	assert.Equal(len(phones), testConfig.ExpectedCount)

	t.Cleanup(func() {
		test_helpers.DeletePhoneNumbersByCountryCode(db, testConfig.CountryCode)
		t.Log(fmt.Sprintf("Performing %s numbers clean up...", testConfig.CountryCode))
	})
}

func TestRussianPhoneNumberValidation(t *testing.T) {
	testConfig := PhoneNumberValidationTestConfig{
		URL:           "https://mysmsbox.ru/",
		Regex:         RuRegex,
		CountryCode:   "RU",
		ExpectedCount: 37,
		WebPageId:     1,
	}
	testPhoneNumberValidation(t, testConfig)
}

func TestUSAPhoneNumberValidation(t *testing.T) {
	testConfig := PhoneNumberValidationTestConfig{
		URL:           "https://www.thisnumber.com/270-258",
		Regex:         USRegex,
		CountryCode:   "US",
		ExpectedCount: 151,
		WebPageId:     1,
	}
	testPhoneNumberValidation(t, testConfig)
}

// There are 15-16 duplicate numbers in the html, so we store only 15
// // DE = Germany
// // 5 Duplicates, only 1 Correct

func TestDEPhoneNumberValidation(t *testing.T) {
	testConfig := PhoneNumberValidationTestConfig{
		URL:           "https://allaboutberlin.com/guides/dial-phone-numbers-germany",
		Regex:         DERegex,
		CountryCode:   "DE",
		ExpectedCount: 1,
		WebPageId:     1,
	}
	testPhoneNumberValidation(t, testConfig)
}

// // FIXME: the nl regexp validation is a bit wrong but it works for now

func TestConcurrentPhoneProcessing(t *testing.T) {
	html, err := ReadFilesInDirectory("test_data")
	if err != nil {
		t.Fatal(err)
	}
	phoneProcessor := NewPhoneProcessor()

	testCountryRegexPatterns := map[string]string{
		"NL": NLRegex,
		"RU": RuRegex,
	}

	phoneProcessor.ProcessPhoneNumbers(html, 1, testCountryRegexPatterns)
	t.Cleanup(func() {
		for countryCode, _ := range testCountryRegexPatterns {
			logger.Debugf("removing %s phones from test database", countryCode)
			test_helpers.DeletePhoneNumbersByCountryCode(db, countryCode)

		}
	})

}
func TestRegexpForTelTag(t *testing.T) {
	regex := `<(?:a|input)[^>]*(?:\s+href\s*=\s*["']\s*tel:([^"']*)["']|type\s*=\s*["']\s*tel["'][^>]*)>`

	html, err := fsutils.ReadTextFile("test_data/tel.html")
	if err != nil {
		t.Fatal(err)
	}
	validator, err := NewPhoneNumberValidator(regex, "")
	phones := validator.FindPhoneNumbers(html)
	phoneHTML := strings.Join(phones, " ")
	phoneProcessor := NewPhoneProcessor()
	testCountryRegexPatterns := map[string]string{
		"NL": NLRegex,
		"RU": RuRegex,
		"US": USRegex,
		"DE": DERegex,
	}

	phoneProcessor.ProcessPhoneNumbers(phoneHTML, 1, testCountryRegexPatterns)
	t.Cleanup(func() {
		for countryCode, _ := range testCountryRegexPatterns {
			logger.Debugf("removing %s phones from test database", countryCode)
			test_helpers.DeletePhoneNumbersByCountryCode(db, countryCode)

		}
	})
}

func TestTearDown(t *testing.T) {
	result := db.Exec("DELETE FROM web_pages WHERE id = ?", 1)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
}
