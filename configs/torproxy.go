package configs

type TorConfig struct {
	//Host is the ip adress that tor is running in. localhost is used by default
	Host string `mapstructure:"host"`
	//Port is the port where tor is running. 9050 is the default
	Port string `mapstructure:"port"`
}

func setupTor() {
	var config TorConfig
	loadConfig("tor", &config)

	cfg.Tor = config
}
