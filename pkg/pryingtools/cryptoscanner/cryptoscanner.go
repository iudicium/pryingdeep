package cryptoscanner

import (
	"regexp"
	"sync"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/lib/pq"

	"github.com/iudicium/pryingdeep/models"
	"github.com/iudicium/pryingdeep/pkg/logger"
)

type CryptoPattern struct {
	Name     string
	Patterns []string
	Results  *pq.StringArray
}

type CryptoScanner struct {
	Crypto         *models.Crypto
	CryptoPatterns []CryptoPattern
}

func NewCryptoScanner() *CryptoScanner {
	crypto := models.Crypto{}

	cryptoPatterns := []CryptoPattern{
		{
			Name: "Wallets",
			Patterns: []string{
				//BTC
				`\b(bc1[ac-hj-np-z02-9]{8,87}|bc[13][a-km-zA-HJ-NP-Z1-9]{25,35})\b`,
				//ETH
				`\b0x[a-f0-9]{40}\b`,
				//XMR - (sometimes fails)
				//https://monero.stackexchange.com/questions/1601/how-to-perform-a-simple-verification-of-a-monero-address-with-a-regular-expressi
				`[48][0-9AB][1-9A-HJ-NP-Za-km-z]{93}`,
				`^(?:[48][0-9AB]|4[1-9A-HJ-NP-Za-km-z]{12}(?:[1-9A-HJ-NP-Za-km-z]{30})?)[1-9A-HJ-NP-Za-km-z]{93}$`,
			},

			Results: &crypto.Wallets,
		},
		{
			Name:     "PGPKeys",
			Patterns: []string{`(?s)-----BEGIN PGP PUBLIC KEY BLOCK-----\n(.*?)\n-----END PGP PUBLIC KEY BLOCK-----`},
			Results:  &crypto.PGPKeys,
		},
		{
			Name:     "Certificates",
			Patterns: []string{`(?s)-----BEGIN CERTIFICATE-----\n(.*?)\n-----END CERTIFICATE-----`},
			Results:  &crypto.Certificates,
		},
	}

	return &CryptoScanner{Crypto: &crypto, CryptoPatterns: cryptoPatterns}
}
func (s *CryptoScanner) searchWithPatterns(html string, patterns []string, cryptoField *pq.StringArray) {
	uniqueMatches := mapset.NewSet[string]()

	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindAllString(html, -1)

		for _, match := range matches {
			uniqueMatches.Add(match)
		}
	}

	*cryptoField = uniqueMatches.ToSlice()
}

// Search spawns multiple goroutines to search for multiple crypto patterns.
func (s *CryptoScanner) Search(html string, pageId int) {
	var wg sync.WaitGroup
	wg.Add(len(s.CryptoPatterns))

	for i := range s.CryptoPatterns {
		go func(p *CryptoPattern) {
			defer wg.Done()
			s.searchWithPatterns(html, p.Patterns, p.Results)
		}(&s.CryptoPatterns[i])
	}

	wg.Wait()
	if s.anyPatternMatched() {
		logger.Infof("Crypto record successfully created on pageID: %d", pageId)
		s.Crypto.WebPageID = pageId
		_, err := models.CryptoCreate(*s.Crypto)
		if err != nil {
			logger.Errorf("Creation of crypto was unsuccessful. %s", err)
		}
	}
}

// anyPatternMatched looks for any matches that have been found and returns a bool value.
func (s *CryptoScanner) anyPatternMatched() bool {
	for _, pattern := range s.CryptoPatterns {
		if len(*pattern.Results) > 0 {
			return true
		}
	}
	return false
}
