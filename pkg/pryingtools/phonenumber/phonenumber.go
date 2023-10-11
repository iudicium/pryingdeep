package phonenumber

import (
	"fmt"
	"github.com/nyaruka/phonenumbers"
	"regexp"
)

// TODO: it looks like i will have to create regex patterns for each country.
// TODO: Let's add support for 5 countries.
// TODO: Russia, USA, UK, Germany, Netherlands
func FindPhoneNumbers(html string) {
	pattern := `\+?[0-9][0-9()\s.-]{8,20}\d`

	regex := regexp.MustCompile(pattern)

	matches := regex.FindAllString(html, -1)
	for i, match := range matches {
		fmt.Println(i, match)

		phoneNumber, err := phonenumbers.Parse(match, "US")
		if err != nil {
			fmt.Printf("Error parsing phone number: %s\n", err)
			continue // Skip to the next match if there's an error
		}

		// You can access information from phoneNumber object here, e.g., countryCode, nationalNumber, etc.
		fmt.Printf("Country Code: %d, National Number: %d\n", phoneNumber.GetCountryCode(), phoneNumber.GetNationalNumber())
	}
}
