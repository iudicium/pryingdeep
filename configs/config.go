package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"



)


type Socks5Config struct {
	Host string
	Port string
}

type DBConfig struct {
	Host string
	Port string
	DbName string
	User string
	Password string
	DbURL string
}

// TODO: probably need  a better name
type Configuration struct {
	TorConf Socks5Config
	DbConf DBConfig

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

func SetupDatabase() DBConfig {
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

    config := DBConfig{
        Host:     DBHost,
        Port:     DBPort,
        DbName:   DbName,
        User:     DBUser,
        Password: DBPass,
        DbURL:    dbURL,
    }

    return config
}




func SetupEnvironment() Configuration {
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  Socks5Config :=  SetupSocks5s()
  DbConfig :=  SetupDatabase()

  confing := Configuration{
		TorConf: Socks5Config,
		DbConf: DbConfig,
  } 
  return confing 
}