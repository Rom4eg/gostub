package config

import (
	"os"
	"path/filepath"
)

func FindConfig() (string, error) {
	l := []string{}

	if uhd, err := os.UserHomeDir(); err == nil {
		l = append(l, filepath.Join(uhd, ".config/gostub/config.yaml"))
	}

	l = append(l, "/etc/gostub/config.yaml")

	if abs, err := filepath.Abs("./config.yaml"); err == nil {
		l = append(l, abs)
	}

	for _, v := range l {
		if _, err := os.Stat(v); err == nil {
			return v, nil
		}
	}

	return "", ErrConfigNotFound
}
