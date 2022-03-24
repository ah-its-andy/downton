package logging

type CommonLogger struct {
	loggers []SimpleLogger
}

func (clg *CommonLogger) Log(level string, message string, args ...interface{}) {
	for _, logger := range clg.loggers {
		logger.Log(level, message, args...)
	}
}
func (clg *CommonLogger) Debug(args ...interface{}) {
	clg.Log(LevelDebug, "", args...)
}
func (clg *CommonLogger) Debugf(format string, args ...interface{}) {
	clg.Log(LevelDebug, format, args...)
}
func (clg *CommonLogger) Info(args ...interface{}) {
	clg.Log(LevelInfo, "", args...)
}
func (clg *CommonLogger) Infof(format string, args ...interface{}) {
	clg.Log(LevelInfo, format, args...)
}
func (clg *CommonLogger) Warn(args ...interface{}) {
	clg.Log(LevelWarn, "", args...)
}
func (clg *CommonLogger) Warnf(format string, args ...interface{}) {
	clg.Log(LevelWarn, format, args...)
}
func (clg *CommonLogger) Error(args ...interface{}) {
	clg.Log(LevelError, "", args...)
}
func (clg *CommonLogger) Errorf(format string, args ...interface{}) {
	clg.Log(LevelError, format, args...)
}
func (clg *CommonLogger) Fatal(args ...interface{}) {
	clg.Log(LevelFatal, "", args...)
}
func (clg *CommonLogger) Fatalf(format string, args ...interface{}) {
	clg.Log(LevelFatal, format, args...)
}
