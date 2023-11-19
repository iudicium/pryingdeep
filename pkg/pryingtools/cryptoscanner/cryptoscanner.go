package cryptoscanner

import (
	"regexp"
	"sync"

	"github.com/lib/pq"

	"github.com/iudicium/pryingdeep/models"
	"github.com/iudicium/pryingdeep/pkg/logger"
)

type CryptoScanner struct {
	Crypto models.Crypto
}

func New() *CryptoScanner {
	return &CryptoScanner{}
}

func (p *CryptoScanner) searchWithPattern(html string, pattern string, cryptoField *pq.StringArray) {
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(html, -1)
	if len(matches) != 0 {
		*cryptoField = matches
	}
}

// Search spawns multiple goroutines to search for mulitple  crypto patterns,
func (p *CryptoScanner) Search(html string, pageId int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		p.searchWithPattern(html, `(?s)-----BEGIN PGP PUBLIC KEY BLOCK-----\n(.*?)\n-----END PGP PUBLIC KEY BLOCK-----`, &p.Crypto.PGPKeys)
	}()

	go func() {
		defer wg.Done()
		p.searchWithPattern(html, `(?s)-----BEGIN CERTIFICATE-----\n(.*?)\n-----END CERTIFICATE-----`, &p.Crypto.Certificates)
	}()

	wg.Wait()
	//TODO add proper error handling
	if len(p.Crypto.PGPKeys) != 0 || len(p.Crypto.Certificates) != 0 {
		logger.Infof("[+] Creating a crypto record...")
		p.Crypto.WebPageID = pageId
		models.CryptoCreate(p.Crypto)
	}

}
