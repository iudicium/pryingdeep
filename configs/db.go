package configs

import (
	"fmt"
)

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"pass"`
	TestName string `mapstructure:"testing_name"`
	URL      string
}

func LoadDatabase() {
	var db Database
	loadConfig("database", &db)
	if db.Port == "" {
		db.Port = "5432"
	}
	if db.Host == "" {
		db.Host = "localhost"
	}
	if db.Name == "" {
		db.Name = "postgres"
	}
	if db.User == "" {
		db.User = "postgres"
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db.User, db.Password, db.Host, db.Port, db.Name)
	db.URL = dbURL
	cfg.DB = db
}
