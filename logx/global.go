package logx

func Debug(args ...interface{}) {
	gLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	gLogger.Debugf(template, args...)
}

func Debugw(msg string, keyAndValues ...interface{}) {
	gLogger.Debugw(msg, keyAndValues...)
}

func Info(args ...interface{}) {
	gLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	gLogger.Infof(template, args...)
}

func Infow(msg string, keyAndValues ...interface{}) {
	gLogger.Infow(msg, keyAndValues...)
}

func Warn(args ...interface{}) {
	gLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	gLogger.Warnf(template, args...)
}

func Warnw(msg string, keyAndValues ...interface{}) {
	gLogger.Warnw(msg, keyAndValues...)
}

func Error(args ...interface{}) {
	gLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	gLogger.Errorf(template, args...)
}

func Errorw(msg string, keyAndValues ...interface{}) {
	gLogger.Errorw(msg, keyAndValues...)
}

func Panic(args ...interface{}) {
	gLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	gLogger.Panicf(template, args...)
}

func Panicw(msg string, keyAndValues ...interface{}) {
	gLogger.Panicw(msg, keyAndValues...)
}

func Fatal(args ...interface{}) {
	gLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	gLogger.Fatalf(template, args...)
}

func Fatalw(msg string, keyAndValues ...interface{}) {
	gLogger.Fatalw(msg, keyAndValues...)
}
