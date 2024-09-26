package loggerconfig

import "github.com/kelseyhightower/envconfig"

type Config struct {
	// Level: info, debug, warning, error, fatal, panic.
	Level string `envconfig:"LOGGER_LEVEL"  default:"debug"`
	// Output: /dev/null/, stdout, stderr (logs will store in the {projectRoot}/var/log dir.).
	Output string `envconfig:"LOGGER_OUTPUT" default:"stdout"`
	// Formatter: text, json.
	Formatter string `envconfig:"LOGGER_FORMAT" default:"json"`
	// LogsDir is any dir from root project dir.
	LogsDir string `envconfig:"LOGGER_LOGS_DIR" default:"var/log"`
	// ContextExtraFields determines which fields must be extract from
	// context.Context and passed into log record (see more into ctxenum package).
	ContextExtraFields []string `envconfig:"LOGGER_CONTEXT_EXTRA_FIELD"`
}

func Load() (*Config, error) {
	cfg := new(Config)
	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *Config) GetLoggerLevel() string {
	return c.Level
}

func (c *Config) GetLoggerOutput() string {
	return c.Output
}

func (c *Config) GetLoggerFormatter() string {
	return c.Formatter
}

func (c *Config) GetLoggerLogsDir() string {
	return c.LogsDir
}

func (c *Config) GetLoggerContextExtraFields() []string {
	return c.ContextExtraFields
}
