package phonenumber

import "fmt"

func MapCountryCodeToRegex(countryCodes []string) map[string]string {
	countryMap := make(map[string]string)
	for _, code := range countryCodes {
		switch code {
		case "DE":
			countryMap[code] = DERegex
		case "RU":
			countryMap[code] = RuRegex
		case "US":
			countryMap[code] = USRegex
		case "GB":
			countryMap[code] = UKRegex
		case "NL":
			countryMap[code] = NLRegex

		}
	}
	fmt.Println(countryMap)
	return countryMap
}
