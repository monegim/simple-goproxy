package main

type Config struct {
	Addr string
}

func main() {
	config := loadConfig()
}

func loadConfig() *Config {
	config := &Config{}
	return config
}
