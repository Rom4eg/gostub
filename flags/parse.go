package flags

import "flag"

func Parse(args []string) {
	if f == nil {
		f = new(Flags)
	}

	flag.StringVar(&f.configFile, "config", "./config.yaml", "config file")
	flag.StringVar(&f.logLevel, "logging", "info", "logging level")

	err := flag.CommandLine.Parse(args)
	if err != nil {
		panic(err)
	}
}
