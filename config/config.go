package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Telegram struct {
		Token string `yaml:"token"`
	}
	Server struct {
		Addr string `yaml:"addr"`
	}

	Database struct {
		path string `yaml:"path"`
	}
}

func MustLoadConfig(path string) (*Config, error) {
	var cfg Config
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
