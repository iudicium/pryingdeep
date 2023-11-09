package email

import (
	"regexp"
)

// EmailFinder performs simple regexp searching.
// Currently, there's no need for validation, but we can definitely take a look at implementing that in the future
type EmailFinder struct {
	emailRegex *regexp.Regexp
}

func NewEmailFinder() *EmailFinder {
	return &EmailFinder{
		emailRegex: regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,3}\b`),
	}
}

func (ef *EmailFinder) FindEmails(html string) []string {
	return ef.emailRegex.FindAllString(html, -1)
}
