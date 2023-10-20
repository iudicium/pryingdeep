package configs

import (
	"fmt"
	"os"
)

type DBConfig struct {
	Host       string
	Port       string
	DbName     string
	User       string
	Password   string
	DbURL      string
	DbTestName string
}

func LoadDatabase() {
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

	DBTestingName := os.Getenv("DB_TESTING_NAME")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DBUser, DBPass, DBHost, DBPort, DbName)

	cfg.DbConf = DBConfig{
		Host:       DBHost,
		Port:       DBPort,
		DbName:     DbName,
		User:       DBUser,
		Password:   DBPass,
		DbURL:      dbURL,
		DbTestName: DBTestingName,
	}
}
