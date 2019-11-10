package logs

const (
	levelInfo = iota
	levelWarn
	levelErr
)

type Logger interface {
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Err(format string, args ...interface{})
}

type stdoutlogger struct{}
