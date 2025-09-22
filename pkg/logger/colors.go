package flogger

type LoggerColor string

const (
	Reset   LoggerColor = "\033[0m"
	Black   LoggerColor = "\033[30m"
	Red     LoggerColor = "\033[31m"
	Green   LoggerColor = "\033[32m"
	Yellow  LoggerColor = "\033[33m"
	Blue    LoggerColor = "\033[34m"
	Magenta LoggerColor = "\033[35m"
	Cyan    LoggerColor = "\033[36m"
	Gray    LoggerColor = "\033[37m"
	White   LoggerColor = "\033[97m"
)

type LogLevel string

const (
	DebugLvl LogLevel = "Debug"
	InfoLvl  LogLevel = "Info"
	WarnLvl  LogLevel = "Warn"
	ErrorLvl LogLevel = "Error"
	FatalLvl LogLevel = "Fatal"
)
