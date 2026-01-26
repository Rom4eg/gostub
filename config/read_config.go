package config

import (
	"bytes"
	"os"

	"gopkg.in/yaml.v2"
)

func _readConfig() (*Config, error) {
	cf := GetConfigLocation()
	if _, err := os.Stat(cf); err != nil {
		return nil, err
	}

	ch, err := os.Open(cf)
	if err != nil {
		return nil, err
	}
	defer ch.Close()

	var rawCfg bytes.Buffer
	n, err := rawCfg.ReadFrom(ch)
	if err != nil {
		return nil, err
	}

	if n == 0 {
		return nil, ErrEmptyConfig
	}

	c := new(Config)
	err = yaml.Unmarshal(rawCfg.Bytes(), c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
