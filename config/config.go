package config

import "os"

type Config struct {
	Domain string
}

func New() *Config {
	return &Config{
		Domain: getEnv("DOMAIN", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
