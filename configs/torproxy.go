package configs

import (
	"log"
	"os"
)

type TorConfig struct {
	Host string
	Port string
}

func SetupTor() {
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
