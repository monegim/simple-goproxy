package main

import (
	"flag"
	rawLog "log"
	"os"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	version bool

	Debug int
	Addr  string
}

func main() {
	config := loadConfig()

	if config.Debug != 0 {
		rawLog.SetFlags(rawLog.LstdFlags | rawLog.Lshortfile)
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	if config.Debug == 2 {
		log.SetReportCaller(true)
	}
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

func loadConfig() *Config {
	config := loadConfigFromCli()
	return config
}

func loadConfigFromCli() *Config {
	config := new(Config)

	flag.BoolVar(&config.version, "version", false, "show simple-goproxy version")
	flag.IntVar(&config.Debug, "debug", 0, "debug mode: 1 - print debug log, 2 - show debug from")
	flag.StringVar(&config.Addr, "addr", ":9080", "proxy listen addr")
	flag.Parse()

	return config
}
