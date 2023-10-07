package tests

import (
	"github.com/r00tk3y/prying-deep/pkg/pryingtools/email"
	"os"
	"path/filepath"
	"testing"
)

var filePath string
var fileContents []byte

func init() {
	filePath = filepath.Join("data", "email.html")

	var err error
	fileContents, err = os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
}

func TestEmail(t *testing.T) {
	HtmlString := string(fileContents)

	matches := email.FindEmail(HtmlString)

	if len(matches) != 3 {
		t.Errorf("expected array length should be 3")
	}
}
func TestNoEmail(t *testing.T) {
	filePath = filepath.Join("data", "no_email.html")

	fileContents, _ = os.ReadFile(filePath)

	matches := email.FindEmail(string(fileContents))
	if len(matches) != 0 {
		t.Errorf("expected array length should be 1")
	}

}
