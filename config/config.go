package config

import "os"

type Config struct {
	Domain string
	Email  string
}

func New() *Config {
	return &Config{
		Domain: getEnv("DOMAIN", ""),
		Email:  getEnv("EMAIL", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
