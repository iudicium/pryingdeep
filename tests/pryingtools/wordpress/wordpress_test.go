package tests

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/iudicium/pryingdeep/pkg/pryingtools/wordpress"
)

var client *http.Client
var url string
var wpFinder *wordpress.WordpressFinder

func TestSetup(t *testing.T) {
	client = &http.Client{}
	wpFinder = wordpress.NewWordpressPatternFinder()

}
func TestWordpressPatternsInHtml(t *testing.T) {
	assert := assert.New(t)
	url = "https://based.win/"
	resp, err := client.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	matches := wpFinder.Find(string(body))

	assert.Equal(len(matches), 127)
}

func TestNoWordpressPatternsInHtml(t *testing.T) {
	assert := assert.New(t)
	url = "https://example.com/"
	resp, err := client.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	matches := wpFinder.Find(string(body))

	assert.Equal(len(matches), 0)
}
