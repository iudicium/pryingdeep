package email

import (
	"regexp"

	mapset "github.com/deckarep/golang-set/v2"
)

// EmailFinder performs simple regexp email searching.
// Currently, there's no need for validation, but we can definitely take a look at implementing that in the future
type EmailFinder struct {
	emailRegex *regexp.Regexp
}

func NewEmailFinder() *EmailFinder {
	return &EmailFinder{
		emailRegex: regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,3}\b`),
	}
}

// FindEmails - just a function that finds unique emails in a webpage. There's no reason to keep the duplicates
func (ef *EmailFinder) FindEmails(html string) []string {
	uniqueMatches := mapset.NewSet[string]()

	matches := ef.emailRegex.FindAllString(html, -1)
	if len(matches) > 0 {
		for _, match := range matches {
			uniqueMatches.Add(match)
		}
	}
	return uniqueMatches.ToSlice()

}
