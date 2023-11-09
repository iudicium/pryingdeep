package configs

type LoggerConfig struct {
	//Level is a simple logger level
	Level string `mapstructure:"level"`
	//Encoder is the encoder that we are using for the logger. E.G json
	Encoder string `mapstructure:"encoder"`
}

func setupLogger() {
	var config LoggerConfig
	loadConfig("logger", &config)
	cfg.Logger = config
}
