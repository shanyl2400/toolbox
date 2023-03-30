package config

import (
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type Config struct {
	BoltDBPath string `yaml:"boltdb_path"`

	HttpPort   int    `yaml:"http_port"`
	AssetsPath string `yaml:"assets_path"`
}

var (
	_config     *Config
	_configOnce sync.Once
)

func Set(c *Config) {
	_config = c
}

func Get() *Config {
	_configOnce.Do(func() {
		_config = &Config{
			HttpPort:   8088,
			BoltDBPath: "./data.db",
			AssetsPath: "./assets",
		}
		conf, err := os.ReadFile("./config.yaml")
		if err != nil {
			log.Println("read config file failed")
			return
		}
		err = yaml.Unmarshal(conf, _config)
		if err != nil {
			log.Println("bad config file format")
			return
		}
	})
	return _config
}
