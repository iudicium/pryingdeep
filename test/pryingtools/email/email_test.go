package tests

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/iudicium/pryingdeep/configs"
	"github.com/iudicium/pryingdeep/models"
	"github.com/iudicium/pryingdeep/pkg/logger"
	"github.com/iudicium/pryingdeep/pkg/pryingtools/email"
)

var client *http.Client

func TestSetup(t *testing.T) {
	configs.SetupEnvironment()
	cfg := configs.GetConfig().DB

	logger.InitLogger(false)
	defer logger.Logger.Sync()

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.TestName)
	models.SetupDatabase(dbURL)
	client = &http.Client{}

}

func TestEmailInHtml(t *testing.T) {
	emailFinder := email.NewEmailFinder()
	testCases := []struct {
		url         string
		expected    int
		description string
	}{
		{
			url:         "https://proton.me/support/zendesk",
			expected:    13,
			description: "Valid HTML with email addresses",
		},
		{
			url:         "https://example.com/",
			expected:    0,
			description: "Valid HTML with no email addresses",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			assert := assert.New(t)

			resp, err := http.Get(tc.url)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)

			matches := emailFinder.FindEmails(string(body))
			assert.Equal(len(matches), tc.expected)
		})
	}
}
