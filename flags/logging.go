package flags

func (f *Flags) Logging() string {
	return f.logLevel
}

func (f *Flags) HasLogging() bool {
	return f.logLevel != ""
}
