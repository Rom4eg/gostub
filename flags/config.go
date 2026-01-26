package flags

func (f *Flags) Config() string {
	return f.configFile
}

func (f *Flags) HasConfig() bool {
	return f.configFile != ""
}
