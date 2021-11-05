package zap

import "go.uber.org/zap"

const (
	EMERGENCY="Emergency"
	ALERT="Alert"
	CRITICAL="Critical"
	ERROR="Error"
	WARNING="Warning"
	NOTICE="Notice"
	INFO="Informational"
	DEBUG="Debug"
	SYSTEM="System"
)

type Logger interface {
	Info(message string)
	Error(message string)
	Debug(message string)
	Warning(message string)
}

var logger Logger
var zaplog *zap.Logger