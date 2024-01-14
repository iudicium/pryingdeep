package email

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var client *http.Client

func TestSetup(t *testing.T) {
	client = &http.Client{}
}

func TestEmailInHtml(t *testing.T) {
	emailFinder := NewEmailFinder()
	testCases := []struct {
		url         string
		expected    int
		description string
	}{
		{
			url:         "https://proton.me/support/zendesk",
			expected:    5,
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
