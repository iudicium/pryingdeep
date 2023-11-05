package configs

import (
	"log"
	"os"
)

type LoggerConfig struct {
	//Level is a simple logger level
	Level string
	//Encoder is the encoder that we are using for the logger. E.G json
	Encoder string
}

func setupLogger() {
	LoggerLevel := os.Getenv("LOGGER_LEVEL")
	if LoggerLevel == "" {
		log.Fatal("Log Level was not specified")
	}
	Encoder := os.Getenv("LOGGER_ENCODER")

	cfg.LoggerConf = LoggerConfig{
		Level:   LoggerLevel,
		Encoder: Encoder,
	}
}
