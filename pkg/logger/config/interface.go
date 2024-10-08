package loggerconfig

type Configurator interface {
	GetLoggerLevel() string
	GetLoggerOutput() string
	GetLoggerFormatter() string
	GetLoggerLogsDir() string
	GetLoggerContextExtraFields() []string
}
