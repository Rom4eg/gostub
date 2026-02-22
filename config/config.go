package config

type Config struct {
	Services []Service `yaml:"services"`
}

var cfg *Config
