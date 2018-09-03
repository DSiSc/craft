package log

import (
	"io"
	"os"
	"time"
)

// Appender is responsible for delivering LogEvents to their destination.
type Appender struct {
	LogLevel      Level
	Output        io.Writer
	Format        string
	ShowCaller    bool
	ShowTimestamp bool
}

// OutputFlags are printed in log record that can be customized.
type OutputFlags struct {
	// TimestampFieldName is the field name used for the timestamp field.
	TimestampFieldName string
	// LevelFieldName is the field name used for the level field.
	LevelFieldName string
	// MessageFieldName is the field name used for the message field.
	MessageFieldName string
	// ErrorFieldName is the field name used for error fields.
	ErrorFieldName string
	// CallerFieldName is the field name used for caller field.
	CallerFieldName string
}

// Config includes configurations for our log, such as log-level.
// For more log destinations just add "Appender" into "Config.[]Appenders".
type Config struct {
	Enabled         bool
	Provider        Provider
	GlobalLogLevel  Level
	TimeStampFormat string
	Appenders       []Appender
	OutputFlags     *OutputFlags
}

// stdoutAppender is a pre-configed console log.
var stdoutAppender = &Appender{
	LogLevel:      DebugLevel,
	Output:        os.Stdout,
	Format:        TextFmt,
	ShowCaller:    true,
	ShowTimestamp: true,
}

// globalOutputFlags contains pre-defined output flags. Usually no need to modify.
var globalOutputFlags = &OutputFlags{
	TimestampFieldName: "time",
	LevelFieldName:     "level",
	MessageFieldName:   "message",
	ErrorFieldName:     "error",
	CallerFieldName:    "caller",
}

// globalConfig is a set of default log configuration with only one "stdoutAppender".
var globalConfig = &Config{
	Enabled:         true,
	Provider:        Zerolog,
	GlobalLogLevel:  DebugLevel,
	TimeStampFormat: time.RFC3339,
	Appenders:       []Appender{*stdoutAppender},
	OutputFlags:     globalOutputFlags,
}

// buildLogger builds a "Logger" with a number of backend logger inside.
// Each logger corresponds to an "Appender".
func (config *Config) buildLogger() Logger {
	if !config.Enabled {
		return nonLogger{}
	}
	switch config.Provider {
	case Zerolog:
		logger := buildZeroLogger(config)
		return logger
	}
	return nil
}

// Level defines log levels.
type Level uint8

const (
	// DebugLevel defines debug log level.
	DebugLevel Level = iota
	// InfoLevel defines info log level.
	InfoLevel
	// WarnLevel defines warn log level.
	WarnLevel
	// ErrorLevel defines error log level.
	ErrorLevel
	// FatalLevel defines fatal log level.
	FatalLevel
	// PanicLevel defines panic log level.
	PanicLevel
	// Disabled disables the logger.
	Disabled
)

// Provider enumerates backend log libs.
type Provider uint8

const (
	Zerolog Provider = iota
)

const (
	// JsonFmt indicates that log output generated in form of JSON.
	JsonFmt string = "JSON"
	// TextFmt indicates that log output generated in form of TEXT.
	TextFmt string = "TEXT"
)
