package phonenumber

// MapCountryCodeToRegex maps each countryCode to the specific regexp
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
		case "NL":
			countryMap[code] = NLRegex

		}
	}
	return countryMap
}
