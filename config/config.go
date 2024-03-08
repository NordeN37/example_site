package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	TLSConfig
}

type TLSConfig struct {
	CreateTLSConfig bool   `envconfig:"CREATE_TLS_CONFIG" default:"false"` // create TLSConfig and run https
	Domain          string `envconfig:"DOMAIN" default:"example.com"`      // domain example: example.com
	Email           string `envconfig:"EMAIL" default:"example@gmail.com"` // email
}

func New() *Config {
	cfg := Config{}
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("failed to load envconfig, err: %s", err)
	}
	return &cfg
}
