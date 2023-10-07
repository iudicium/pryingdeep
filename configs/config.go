package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type TorConfig struct {
	Host string
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	DbName   string
	User     string
	Password string
	DbURL    string
}

type Configuration struct {
	TorConf    TorConfig
	DbConf     DBConfig
	LoggerConf LoggerConfig
}

type LoggerConfig struct {
	Level   string
	Encoder string
}

var cfg Configuration

func GetConfig() *Configuration {
	return &cfg
}

func SetupLogger() {
	LoggerLevel := os.Getenv("LOGGER_LEVEL")
	if LoggerLevel == "" {
		log.Fatal("Log Level was not specified")
	}
	Encoder := os.Getenv("LOGGER_ENCODER")
	if Encoder == "" {
		log.Fatal("Logger encoder was not specified")
	}
	cfg.LoggerConf = LoggerConfig{
		Level:   LoggerLevel,
		Encoder: Encoder,
	}

}
func SetupSocks5s() {
	Socks5Port := os.Getenv("SOCKS5_PORT")
	if Socks5Port == "" {
		log.Fatal("Socks5 port was not specified")
	}

	Socks5Host := os.Getenv("SOCKS5_HOST")
	if Socks5Host == "" {
		log.Fatal("Socks5 host was not specified ")
	}

	cfg.TorConf = TorConfig{
		Host: Socks5Host,
		Port: Socks5Port,
	}

}

func SetupDatabase() {
	DBHost := os.Getenv("DB_HOST")
	if DBHost == "" {
		DBHost = "localhost"
	}

	DBPort := os.Getenv("DB_PORT")
	if DBPort == "" {
		DBPort = "5432"
	}

	DbName := os.Getenv("DB_NAME")
	if DbName == "" {
		DbName = "postgres"
	}

	DBUser := os.Getenv("DB_USER")
	if DBUser == "" {
		DBUser = "postgres"
	}

	DBPass := os.Getenv("DB_PASS")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DBUser, DBPass, DBHost, DBPort, DbName)

	cfg.DbConf = DBConfig{
		Host:     DBHost,
		Port:     DBPort,
		DbName:   DbName,
		User:     DBUser,
		Password: DBPass,
		DbURL:    dbURL,
	}
}

func SetupEnvironment() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	SetupSocks5s()
	SetupDatabase()
	SetupLogger()

}
