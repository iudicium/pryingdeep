package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/iudicium/pryingdeep/pkg/fsutils"
	"github.com/iudicium/pryingdeep/pkg/pryingtools/favicon"
)

var filePath string
var fileContents []byte

func init() {
	filePath = filepath.Join("data", "favicon.html")

	var err error
	fileContents, err = os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
}

func TestFaviconExtraction(t *testing.T) {
	HtmlString := string(fileContents)
	baseUrl := "http://explorerzydxu5ecjrkwceayqybizmpjjznk5izmitf2modhcusuqlid.onion"
	torProxy := "socks5://localhost:9050"
	urls := favicon.ExtractFaviconUrls(HtmlString, baseUrl)
	if len(urls) != 1 {
		t.Errorf("there's only 1 favicon in that baseUrl")
	}

	hashes := favicon.GetFaviconHash(urls, torProxy)
	fmt.Println(hashes)

}

func TestIconHash(t *testing.T) {
	tests := []struct {
		filePath     string
		expectedHash string
	}{
		{"data/favicon.png", "-398941349"},
		{"data/logo.png", "1847410482"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("File: %s", tt.filePath), func(t *testing.T) {
			assert := assert.New(t)

			file, err := fsutils.ReadTextFile(tt.filePath)
			if err != nil {
				t.Fatal(err)
			}

			hash := favicon.IconHash([]byte(file))

			assert.Equal(hash, tt.expectedHash)
		})
	}
}
