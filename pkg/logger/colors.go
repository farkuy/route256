package logger_custome

type LoggerColor string

const (
	Reset     LoggerColor = "\033[0m"
	Black     LoggerColor = "\033[30m"
	Red       LoggerColor = "\033[31m"
	Green     LoggerColor = "\033[32m"
	Yellow    LoggerColor = "\033[33m"
	Blue      LoggerColor = "\033[34m"
	Magenta   LoggerColor = "\033[35m"
	Cyan      LoggerColor = "\033[36m"
	Gray      LoggerColor = "\033[37m"
	White     LoggerColor = "\033[97m"
	Bold      LoggerColor = "\033[1m"
	Italic    LoggerColor = "\033[3m"
	Underline LoggerColor = "\033[4m"
	Invert    LoggerColor = "\033[7m"
)
