package tests

import (
	"fmt"
	"github.com/r00tk3y/prying-deep/configs"
	"github.com/r00tk3y/prying-deep/models"
	"github.com/r00tk3y/prying-deep/pkg/logger"
	"github.com/r00tk3y/prying-deep/pkg/pryingtools/wordpress"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

var client *http.Client
var url string

func TestSetup(t *testing.T) {
	configs.SetupEnvironment()
	cfg := configs.GetConfig().DbConf

	logger.InitLogger()
	defer logger.Logger.Sync()

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbTestName)
	models.SetupDatabase(dbURL)
	client = &http.Client{}

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

	matches, err := wordpress.FindWordpressPatterns(string(body))
	if err != nil {
		t.Error("something went wrong during test of wordpress", err)
	}

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

	matches, err := wordpress.FindWordpressPatterns(string(body))
	if err != nil {
		t.Error("something went wrong during test of wordpress", err)
	}

	assert.Equal(len(matches), 0)
}
