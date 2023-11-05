package configs

// PryingConfig is the modules that the user specifies during crawling.
type PryingConfig struct {
	Email     bool `json:"email"`
	Crypto    bool `json:"crypto"`
	Wordpress bool `json:"wordpress"`
	//PhoneNumbers List of countries. RU,NL,DE,GB,US. You can specify multiple or just one.
	//Default is blank
	PhoneNumbers []string `json:"phoneNumbers"`
}

func loadPryingConfig() {
	var config PryingConfig
	loadConfig("configs/json/pryingConfig.json", &config)
	cfg.PryingConf = config
}
