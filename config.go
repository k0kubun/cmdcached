package main

import (
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

const (
	configFile = ".cmdcached"
)

type Config struct {
	CacheConfigs []CacheConfig `toml:"cache"`
}

type CacheConfig struct {
	Command       string `toml:"command"`
	Subscribe     string `toml:"subscribe"`
	EachDirectory bool   `toml:"each_directory"`
}

func (c *Config) Load() {
	path := filepath.Join(homeDir(), configFile)
	if !fileExists(path) {
		return
	}

	_, err := toml.DecodeFile(path, c)
	if err != nil {
		log.Println(err)
	}
}
