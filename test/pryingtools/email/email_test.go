package tests

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/pryingbytez/pryingdeep/configs"
	"github.com/pryingbytez/pryingdeep/models"
	"github.com/pryingbytez/pryingdeep/pkg/logger"
	"github.com/pryingbytez/pryingdeep/pkg/pryingtools/email"
	"github.com/stretchr/testify/assert"
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

//func TestNoEmail(t *testing.T) {
//	filePath = filepath.Join("data", "no_email.html")
//
//	fileContents, _ = os.ReadFile(filePath)
//
//	matches := email.FindEmail(string(fileContents))
//	if len(matches) != 0 {
//		t.Errorf("expected array length should be 1")
//	}
//
//}

func TestEmailInHtml(t *testing.T) {
	assert := assert.New(t)
	url = "https://proton.me/support/zendesk"
	resp, err := client.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	matches := email.FindEmail(string(body))
	assert.Equal(len(matches), 13)
}
func TestNoEmailInHtml(t *testing.T) {
	assert := assert.New(t)
	url = "https://example.com/"
	resp, err := client.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	matches := email.FindEmail(string(body))
	assert.Equal(len(matches), 0)
}
