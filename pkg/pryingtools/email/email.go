package email

import (
	"regexp"
)

func FindEmail(html string) []string {
	emailRegex := `\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}\b`

	re := regexp.MustCompile(emailRegex)

	matches := re.FindAllString(html, -1)

	return matches
}

//TODO: add support for checking emails in various leaks, websites, etc
//TODO: add tests
