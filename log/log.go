package log

import (
	"io"
	"os"
	"strings"
)

// backendLogger is the actual logging object that we use, and is pre-built as a global variable.
var backendLogger = globalConfig.buildLogger()

// Disable to stop logging.
func Disable() {
	globalConfig.Enabled = false
	backendLogger = &nonLogger{}
}

// Enable to start logging.
func Enable() {
	globalConfig.Enabled = true
	backendLogger = globalConfig.buildLogger()
}

func Debug(msg string) {
	backendLogger.Debug(msg)
}

func Info(msg string) {
	backendLogger.Info(msg)
}

func Warn(msg string) {
	backendLogger.Warn(msg)
}

func Error(msg string) {
	backendLogger.Error(msg)
}

func Fatal(msg string) {
	backendLogger.Fatal(msg)
}

func Panic(msg string) {
	backendLogger.Panic(msg)
}

func DebugKV(msg string, keyvals map[string]interface{}) {
	backendLogger.DebugKV(msg, keyvals)
}

func InfoKV(msg string, keyvals map[string]interface{}) {
	backendLogger.InfoKV(msg, keyvals)
}

func WarnKV(msg string, keyvals map[string]interface{}) {
	backendLogger.WarnKV(msg, keyvals)
}

func ErrorKV(msg string, keyvals map[string]interface{}) {
	backendLogger.ErrorKV(msg, keyvals)
}

func FatalKV(msg string, keyvals map[string]interface{}) {
	backendLogger.FatalKV(msg, keyvals)
}

func PanicKV(msg string, keyvals map[string]interface{}) {
	backendLogger.PanicKV(msg, keyvals)
}

func GetGlobalConfig() *Config {
	return globalConfig
}

// SetGlobalConfig is used to refresh logging manners.
func SetGlobalConfig(config *Config) {
	globalConfig = config
	backendLogger = globalConfig.buildLogger()
}

func GetGlobalLogLevel() Level {
	return globalConfig.GlobalLogLevel
}

// SetGlobalLogLevel is used to restraint log-level of all "Appenders".
func SetGlobalLogLevel(level Level) {
	globalConfig.GlobalLogLevel = level
	backendLogger.SetGlobalLogLevel(level)
}

func GetOutputFlags() *OutputFlags {
	return globalConfig.OutputFlags
}

// SetOutputFlags is used to reconfig output flags.
func SetOutputFlags(flags *OutputFlags) {
	globalConfig.OutputFlags = flags
	backendLogger.SetOutputFlags(flags)
}

func SetTimestampFormat(format string) {
	globalConfig.TimeStampFormat = format
	globalConfig.buildLogger()
}

// AddAppender adds a new logging destination.
func AddAppender(output io.Writer, logLevel Level, format string, showCaller bool, showTimestamp bool) {
	globalConfig.Appenders = append(globalConfig.Appenders, Appender{
		LogLevel:      logLevel,
		Output:        output,
		Format:        format,
		ShowCaller:    showCaller,
		ShowTimestamp: showTimestamp,
	})
	backendLogger = globalConfig.buildLogger()
}

// AddFileAppender adds a new logging destination that append logs to a specified file.
func AddFileAppender(filePath string, logLevel Level, format string, showCaller bool, showTimestamp bool) {
	_, err := os.Stat(filePath)

	if err != nil {
		if os.IsNotExist(err) {
			parentPath := filePath[0:strings.LastIndex(filePath, "/")]
			_, err := os.Stat(filePath)
			if err != nil {
				if os.IsNotExist(err) {
					err := os.MkdirAll(parentPath, 0755)
					if err != nil {
						panic(err)
					}
				}
			}
		} else {
			panic(err)
		}
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	AddAppender(file, logLevel, format, showCaller, showTimestamp)
}

// SetAppenders sets a set of "Appenders".
func SetAppenders(appenders []Appender) {
	globalConfig.Appenders = appenders
	backendLogger = globalConfig.buildLogger()
}
