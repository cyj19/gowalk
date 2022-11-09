package logk

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
	"time"
)

var (
	gLogger Logger
	logName = "gowalk"
	logMu   sync.Mutex
)

type Logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
	Debugw(string, ...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Infow(string, ...interface{})
	Warn(...interface{})
	Warnf(string, ...interface{})
	Warnw(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
	Errorw(string, ...interface{})
	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicw(string, ...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalw(string, ...interface{})
}

type LogConfig struct {
	Level       Level       `json:"level" mapstructure:"level"`
	Mode        Mode        `json:"mode" mapstructure:"mode"`
	TimeFormat  string      `json:"time_format" mapstructure:"time_format"`
	EncodeLevel EncodeLevel `json:"encode_level" mapstructure:"encode_level"`
	Name        string      `json:"name" mapstructure:"name"`
	MaxSize     int         `json:"max_size" mapstructure:"max_size"`
	MaxAge      int         `json:"max_age" mapstructure:"max_age"`
	MaxBackups  int         `json:"max_backups" mapstructure:"max_backups"`
	Compress    bool        `json:"compress" mapstructure:"compress"`
}

type defaultLogger struct {
	zLog *zap.SugaredLogger
}

func (d *defaultLogger) Debug(args ...interface{}) {
	d.zLog.Debug(args...)
}

func (d *defaultLogger) Debugf(template string, args ...interface{}) {
	d.zLog.Debugf(template, args...)
}

func (d *defaultLogger) Debugw(msg string, keyAndValues ...interface{}) {
	d.zLog.Debugw(msg, keyAndValues...)
}

func (d *defaultLogger) Info(args ...interface{}) {
	d.zLog.Info(args...)
}

func (d *defaultLogger) Infof(template string, args ...interface{}) {
	d.zLog.Infof(template, args...)
}

func (d *defaultLogger) Infow(msg string, keyAndValues ...interface{}) {
	d.zLog.Infow(msg, keyAndValues...)
}

func (d *defaultLogger) Warn(args ...interface{}) {
	d.zLog.Warn(args...)
}

func (d *defaultLogger) Warnf(template string, args ...interface{}) {
	d.zLog.Warnf(template, args...)
}

func (d *defaultLogger) Warnw(msg string, keyAndValues ...interface{}) {
	d.zLog.Warnw(msg, keyAndValues...)
}

func (d *defaultLogger) Error(args ...interface{}) {
	d.zLog.Error(args...)
}

func (d *defaultLogger) Errorf(template string, args ...interface{}) {
	d.zLog.Errorf(template, args...)
}

func (d *defaultLogger) Errorw(msg string, keyAndValues ...interface{}) {
	d.zLog.Errorw(msg, keyAndValues...)
}

func (d *defaultLogger) Panic(args ...interface{}) {
	d.zLog.Panic(args...)
}

func (d *defaultLogger) Panicf(template string, args ...interface{}) {
	d.zLog.Panicf(template, args...)
}

func (d *defaultLogger) Panicw(msg string, keyAndValues ...interface{}) {
	d.zLog.Panicw(msg, keyAndValues...)
}

func (d *defaultLogger) Fatal(args ...interface{}) {
	d.zLog.Fatal(args...)
}

func (d *defaultLogger) Fatalf(template string, args ...interface{}) {
	d.zLog.Fatalf(template, args...)
}

func (d *defaultLogger) Fatalw(msg string, keyAndValues ...interface{}) {
	d.zLog.Fatalw(msg, keyAndValues...)
}

var _ Logger = (*defaultLogger)(nil)

// SetupLog 初始化默认日志
func SetupLog(wd string, cf LogConfig) error {

	if cf.Name != "" {
		logName = cf.Name
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	if cf.TimeFormat != "" {
		encoderCfg.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(t.Format(cf.TimeFormat))
		}
	}

	encoderCfg.EncodeLevel = cf.EncodeLevel.switchEncodeLevel()

	encoder := zapcore.NewJSONEncoder(encoderCfg)

	logLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= cf.Level.switchLevel()
	})

	writeSync := zapcore.NewMultiWriteSyncer(cf.Mode.switchWriter(wd, cf)...)
	newCore := zapcore.NewCore(encoder, writeSync, logLevel)

	zl := zap.New(newCore, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()

	gLogger = &defaultLogger{
		zLog: zl,
	}

	return nil

}

// GetLogger 获取日志对象
func GetLogger() Logger {
	return gLogger
}

// SetLogger 设置自定义Logger
func SetLogger(l Logger) {
	logMu.Lock()
	defer logMu.Unlock()
	gLogger = l
}
