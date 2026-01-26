package config

func Reload() {
	cfg = nil

	c, err := _readConfig()
	if err != nil {
		panic(err)
	}

	cfg = c
}
