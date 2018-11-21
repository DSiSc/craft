package log

// Logger interface defines all behaviors of a backendLogger.
type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
	Panic(msg string)

	DebugKV(msg string, keyvals map[string]interface{})
	InfoKV(msg string, keyvals map[string]interface{})
	WarnKV(msg string, keyvals map[string]interface{})
	ErrorKV(msg string, keyvals map[string]interface{})
	FatalKV(msg string, keyvals map[string]interface{})
	PanicKV(msg string, keyvals map[string]interface{})

	SetGlobalLogLevel(level Level)
	SetOutputFlags(flags *OutputFlags)
	SetTimeFieldFormat(format string)
}
