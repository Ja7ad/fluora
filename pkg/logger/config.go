package logger

type Config struct {
	Colorful           bool              `yaml:"colorful"`
	MaxBackups         int               `yaml:"max_backups"`
	RotateLogAfterDays int               `yaml:"rotate_log_after_days"`
	Compress           bool              `yaml:"compress"`
	Targets            []string          `yaml:"targets"`
	Levels             map[string]string `yaml:"levels"`
}

func DefaultConfig() *Config {
	conf := &Config{
		Levels:             make(map[string]string),
		Colorful:           true,
		MaxBackups:         0,
		RotateLogAfterDays: 1,
		Compress:           true,
		Targets:            []string{"console", "file"},
	}

	conf.Levels["default"] = "info"
	conf.Levels["_ai"] = "warn"
	conf.Levels["_service"] = "info"
	conf.Levels["_rest"] = "warn"

	return conf
}

// BasicCheck performs basic checks on the configuration.
func (*Config) BasicCheck() error {
	return nil
}
