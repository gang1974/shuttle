package log

import "time"

const (
	LogTrace = 0
	LogDebug = 1
	LogInfo  = 2
	LogError = 3
)

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

type ILogger interface {
	Trace(...interface{})
	Debug(...interface{})
	Info(...interface{})
	Error(...interface{})
	Tracef(string, ...interface{})
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Errorf(string, ...interface{})
}
