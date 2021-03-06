package log

import "github.com/uber-go/zap"

//CM is a Checked Message like
type CM interface {
	Write(fields ...zap.Field)
	OK() bool
}

//D is a debug logger
func D(logger zap.Logger, message string, callback ...func(l CM)) {
	log(logger, zap.DebugLevel, message, callback...)
}

//I is a info logger
func I(logger zap.Logger, message string, callback ...func(l CM)) {
	log(logger, zap.InfoLevel, message, callback...)
}

//W is a warn logger
func W(logger zap.Logger, message string, callback ...func(l CM)) {
	log(logger, zap.WarnLevel, message, callback...)
}

//E is a error logger
func E(logger zap.Logger, message string, callback ...func(l CM)) {
	log(logger, zap.ErrorLevel, message, callback...)
}

//P is a panic logger
func P(logger zap.Logger, message string, callback ...func(l CM)) {
	log(logger, zap.PanicLevel, message, callback...)
}

//F is a fatal logger
func F(logger zap.Logger, message string, callback ...func(l CM)) {
	log(logger, zap.FatalLevel, message, callback...)
}

func defaultWrite(l CM) {
	l.Write()
}

func log(logger zap.Logger, logLevel zap.Level, message string, callback ...func(l CM)) {
	cb := defaultWrite
	if len(callback) == 1 {
		cb = callback[0]
	}
	if cm := logger.Check(logLevel, message); cm.OK() {
		cb(cm)
	}
}
