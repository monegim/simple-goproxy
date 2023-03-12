package main

import "flag"

type Config struct {
	Addr string
}

func main() {
	config := loadConfig()
}

func loadConfig() *Config {
	config := loadConfigFromCli()
	return config
}

func loadConfigFromCli() *Config {
	config := new(Config)

	flag.StringVar(&config.Addr, "addr", ":9080", "proxy listen addr")
	flag.Parse()

	return config
}
