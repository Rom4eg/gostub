package config

var config_location = ""

func SetConfigLocation(location string) {
	config_location = location
}

func GetConfigLocation() string {
	return config_location
}
