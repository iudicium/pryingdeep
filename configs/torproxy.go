package configs

import (
	"log"
	"os"
)

type TorConfig struct {
	//Host is the ip adress that tor is running in. localhost is used by default
	Host string
	//Port is the port where tor is running. 9050 is the default
	Port string
}

func setupTor() {
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
