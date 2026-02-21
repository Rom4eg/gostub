package config

type Service struct {
	Name    string         `yaml:"name"`
	Type    string         `yaml:"type"`
	Options map[string]any `yaml:"options"`
}
