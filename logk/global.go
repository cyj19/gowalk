package logk

func Debug(args ...interface{}) {
	l := dLogger.Skip(1)
	l.Debug(args...)
	dLogger.PoolPut(l)
}

func Debugf(template string, args ...interface{}) {
	l := dLogger.Skip(1)
	l.Debugf(template, args...)
	dLogger.PoolPut(l)
}

func Debugw(msg string, keyAndValues ...interface{}) {
	l := dLogger.Skip(1)
	l.Debugw(msg, keyAndValues...)
	dLogger.PoolPut(l)
}

func Info(args ...interface{}) {
	l := dLogger.Skip(1)
	l.Info(args...)
	dLogger.PoolPut(l)
}

func Infof(template string, args ...interface{}) {
	l := dLogger.Skip(1)
	l.Infof(template, args...)
	dLogger.PoolPut(l)
}

func Infow(msg string, keyAndValues ...interface{}) {
	l := dLogger.Skip(1)
	l.Infow(msg, keyAndValues...)
	dLogger.PoolPut(l)
}

func Warn(args ...interface{}) {
	l := dLogger.Skip(1)
	l.Warn(args...)
	dLogger.PoolPut(l)
}

func Warnf(template string, args ...interface{}) {
	l := dLogger.Skip(1)
	l.Warnf(template, args...)
	dLogger.PoolPut(l)
}

func Warnw(msg string, keyAndValues ...interface{}) {
	l := dLogger.Skip(1)
	l.Warnw(msg, keyAndValues...)
	dLogger.PoolPut(l)
}

func Error(args ...interface{}) {
	l := dLogger.Skip(1)
	l.Error(args...)
	dLogger.PoolPut(l)
}

func Errorf(template string, args ...interface{}) {
	l := dLogger.Skip(1)
	l.Errorf(template, args...)
	dLogger.PoolPut(l)
}

func Errorw(msg string, keyAndValues ...interface{}) {
	l := dLogger.Skip(1)
	l.Errorw(msg, keyAndValues...)
	dLogger.PoolPut(l)
}

func Panic(args ...interface{}) {
	l := dLogger.Skip(1)
	l.Panic(args...)
	dLogger.PoolPut(l)
}

func Panicf(template string, args ...interface{}) {
	l := dLogger.Skip(1)
	l.Panicf(template, args...)
	dLogger.PoolPut(l)
}

func Panicw(msg string, keyAndValues ...interface{}) {
	l := dLogger.Skip(1)
	l.Panicw(msg, keyAndValues...)
	dLogger.PoolPut(l)
}

func Fatal(args ...interface{}) {
	l := dLogger.Skip(1)
	l.Fatal(args...)
	dLogger.PoolPut(l)
}

func Fatalf(template string, args ...interface{}) {
	l := dLogger.Skip(1)
	l.Fatalf(template, args...)
	dLogger.PoolPut(l)
}

func Fatalw(msg string, keyAndValues ...interface{}) {
	l := dLogger.Skip(1)
	l.Fatalw(msg, keyAndValues...)
	dLogger.PoolPut(l)
}
