package config

func Get() *Config {
	if cfg == nil {
		Reload()
	}
	return cfg
}
