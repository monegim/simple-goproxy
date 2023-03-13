package main

import (
	"flag"
	"fmt"
)

type Config struct {
	version bool

	Addr string
}

func main() {
	config := loadConfig()
	fmt.Println(config.Addr)
}

func loadConfig() *Config {
	config := loadConfigFromCli()
	return config
}

func loadConfigFromCli() *Config {
	config := new(Config)

	flag.BoolVar(&config.version, "version", false, "show simple-goproxy version")
	flag.StringVar(&config.Addr, "addr", ":9080", "proxy listen addr")
	flag.Parse()

	return config
}
