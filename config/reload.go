package config

func Reload() {
	Set(nil)

	c, err := _readConfig()
	if err != nil {
		panic(err)
	}

	Set(c)
}
