package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"



)


type Socks5Config struct {
	Host string
	Port string
}

type PostgresConfig struct {
	Host string
	Port string
	DbName string
	User string
	Password string
	DbURL string
}


type Configuration struct {
	TorConf Socks5Config
	DbConf PostgresConfig

}


func SetupSocks5s() Socks5Config {
  Socks5Port :=  os.Getenv("SOCKS5_PORT")
  if Socks5Port == "" {
	log.Fatal("Socks5 port was not specified")
  }

  Socks5Host := os.Getenv("SOCKS5_HOST")
  if Socks5Host == "" {
	log.Fatal("Socks5 host was not specified ")
  }	
	
  return Socks5Config{
	Host: Socks5Host,
	Port:  Socks5Port,
  }
}

func SetupPostgres() PostgresConfig {
	PostgresHost := os.Getenv("POSTGRES_HOST")
	if PostgresHost == "" {
		PostgresHost = "localhost"
	}

	PostgresPort := os.Getenv("POSTGRES_PORT")
	if PostgresPort == "" {
		PostgresPort = "5432"
	}

	DbName := os.Getenv("POSTGRES_DB")
	if DbName == "" {
		DbName = "postgres"
	}

	PostgresUser := os.Getenv("POSTGRES_USER")
	if PostgresUser == "" {
		PostgresUser = "postgres"
	}

	PostgresPass := os.Getenv("POSTGRES_PASS")

	dbUrl := "postgres://" + PostgresUser + ":" + PostgresPass + "@" + PostgresHost + ":" + PostgresPort + "/" + DbName


	config := PostgresConfig{
		Host:     PostgresHost,
		Port:     PostgresPort,
		DbName:   DbName,
		User:     PostgresUser,
		Password: PostgresPass,
		DbURL:    dbUrl,
	}
		

	return config
}

func SetupEnvironment() Configuration {
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  Socks5Config :=  SetupSocks5s()
  PostgresConfig :=  SetupPostgres()

  confing := Configuration{
		TorConf: Socks5Config,
		DbConf: PostgresConfig,
  } 
  return confing 
}