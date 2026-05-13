package config

import "errors"

var (
	ErrEmptyConfig    = errors.New("empty config")
	ErrConfigNotFound = errors.New("config not found")
)
