package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Port  string `toml:"port"`
	Redis *RedisConfig
}

type RedisConfig struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

func NewConfig(path string) *Config {
	var cfg Config
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		log.Fatalf("cannot read %v: %v", path, err)
	}

	return &cfg
}
