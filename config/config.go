package config

type Config struct {
	Host     string    `yaml:"host"`
	Port     int       `yaml:"port"`
	StubRoot string    `yaml:"stub-root"`
	Services []Service `yaml:"services"`
}

var cfg *Config
