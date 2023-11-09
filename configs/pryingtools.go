package configs

// PryingOptions is the modules that the user specifies during crawling.
type PryingOptions struct {
	Email     bool `mapstructure:"email"`
	Crypto    bool `mapstructure:"crypto"`
	Wordpress bool `mapstructure:"wordpress"`
	//PhoneNumbers List of countries. RU,NL,DE,GB,US. You can specify multiple or just one.
	//Default is blank
	PhoneNumbers []string `mapstructure:"phone-numbers"`
}

func loadPryingConfig() {
	var config PryingOptions
	loadConfig("prying", &config)
	cfg.PryingOptions = config
}
