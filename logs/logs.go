package logs

import (
	"fmt"
	"time"
)

func NewLogger() Logger {
	return &stdoutlogger{}
}

func (sl *stdoutlogger) Info(format string, args ...interface{}) {
	fmt.Println(formatString(levelInfo, format, args...))
}

func (sl *stdoutlogger) Warn(format string, args ...interface{}) {
	fmt.Println(formatString(levelWarn, format, args...))
}

func (sl *stdoutlogger) Err(format string, args ...interface{}) {
	fmt.Println(formatString(levelErr, format, args...))
}

func formatString(level int, format string, args ...interface{}) string {
	now := time.Now()

	var prefix string

	switch level {

	case levelInfo:
		prefix = "[INFO]"
	case levelWarn:
		prefix = "[WARN]"
	case levelErr:
		prefix = "[ERR]"
	}

	timeStr := fmt.Sprintf(
		"[%d-%02d-%02d %02d:%02d]",
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
	)

	return prefix + timeStr + " " + fmt.Sprintf(format, args...)
}
