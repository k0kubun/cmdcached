package server

import (
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"gopkg.in/fsnotify.v1"
)

const (
	configFile = ".cmdcached"
)

var (
	configPath = filepath.Join(homeDir(), configFile)
)

type Config struct {
	CacheConfigs []CacheConfig `toml:"cache"`
	watcher      *fsnotify.Watcher
}

type CacheConfig struct {
	Command       string `toml:"command"`
	Subscribe     string `toml:"subscribe"`
	EachDirectory bool   `toml:"each_directory"`
}

func NewConfig() *Config {
	c := new(Config)
	c.Load()

	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println(err)
		return c
	}

	err = w.Add(configPath)
	if err != nil {
		log.Println(err)
		return c
	}

	c.watcher = w
	go c.watchConfig()

	return c
}

func (c *Config) Load() {
	if !fileExists(configPath) {
		return
	}

	_, err := toml.DecodeFile(configPath, c)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("config reloaded")
}

func (c *Config) watchConfig() {
	for {
		select {
		case ev := <-c.watcher.Events:
			c.Load()

			if ev.Op == fsnotify.Rename {
				c.watcher.Remove(configPath)
				c.watcher.Add(configPath)
			}
		}
	}
}
