package logging

import "github.com/ah-its-andy/downton/core"

type CommonLogger struct {
	loggers []core.SimpleLogger
}

func (clg *CommonLogger) Log(level string, message string, args ...any) {
	for _, logger := range clg.loggers {
		logger.Log(level, message, args...)
	}
}
func (clg *CommonLogger) Debug(args ...any) {
	clg.Log(LevelDebug, "", args...)
}
func (clg *CommonLogger) Debugf(format string, args ...any) {
	clg.Log(LevelDebug, format, args...)
}
func (clg *CommonLogger) Info(args ...any) {
	clg.Log(LevelInfo, "", args...)
}
func (clg *CommonLogger) Infof(format string, args ...any) {
	clg.Log(LevelInfo, format, args...)
}
func (clg *CommonLogger) Warn(args ...any) {
	clg.Log(LevelWarn, "", args...)
}
func (clg *CommonLogger) Warnf(format string, args ...any) {
	clg.Log(LevelWarn, format, args...)
}
func (clg *CommonLogger) Error(args ...any) {
	clg.Log(LevelError, "", args...)
}
func (clg *CommonLogger) Errorf(format string, args ...any) {
	clg.Log(LevelError, format, args...)
}
func (clg *CommonLogger) Fatal(args ...any) {
	clg.Log(LevelFatal, "", args...)
}
func (clg *CommonLogger) Fatalf(format string, args ...any) {
	clg.Log(LevelFatal, format, args...)
}
