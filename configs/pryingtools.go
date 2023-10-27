package configs

type PryingConfig struct {
	Email        bool     `json:"email"`
	Crypto       bool     `json:"crypto"`
	Wordpress    bool     `json:"wordpress"`
	PhoneNumbers []string `json:"phoneNumbers"`
}

func loadPryingConfig() {
	var config PryingConfig
	loadConfig("configs/json/pryingConfig.json", &config)
	cfg.PryingConf = config
}
