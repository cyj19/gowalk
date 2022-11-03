package logx

import (
	"github.com/cyj19/gowalk/core"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var (
	defaultLogger Logger
	logName       = "gowalk"
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

type DefaultLogger struct {
	zLog *zap.SugaredLogger
}

func (d *DefaultLogger) Debug(args ...interface{}) {
	d.zLog.Debug(args...)
}

func (d *DefaultLogger) Debugf(template string, args ...interface{}) {
	d.zLog.Debugf(template, args...)
}

func (d *DefaultLogger) Debugw(msg string, keyAndValues ...interface{}) {
	d.zLog.Debugw(msg, keyAndValues...)
}

func (d *DefaultLogger) Info(args ...interface{}) {
	d.zLog.Info(args...)
}

func (d *DefaultLogger) Infof(template string, args ...interface{}) {
	d.zLog.Infof(template, args...)
}

func (d *DefaultLogger) Infow(msg string, keyAndValues ...interface{}) {
	d.zLog.Infow(msg, keyAndValues...)
}

func (d *DefaultLogger) Warn(args ...interface{}) {
	d.zLog.Warn(args...)
}

func (d *DefaultLogger) Warnf(template string, args ...interface{}) {
	d.zLog.Warnf(template, args...)
}

func (d *DefaultLogger) Warnw(msg string, keyAndValues ...interface{}) {
	d.zLog.Warnw(msg, keyAndValues...)
}

func (d *DefaultLogger) Error(args ...interface{}) {
	d.zLog.Error(args...)
}

func (d *DefaultLogger) Errorf(template string, args ...interface{}) {
	d.zLog.Errorf(template, args...)
}

func (d *DefaultLogger) Errorw(msg string, keyAndValues ...interface{}) {
	d.zLog.Errorw(msg, keyAndValues...)
}

func (d *DefaultLogger) Panic(args ...interface{}) {
	d.zLog.Panic(args...)
}

func (d *DefaultLogger) Panicf(template string, args ...interface{}) {
	d.zLog.Panicf(template, args...)
}

func (d *DefaultLogger) Panicw(msg string, keyAndValues ...interface{}) {
	d.zLog.Panicw(msg, keyAndValues...)
}

func (d *DefaultLogger) Fatal(args ...interface{}) {
	d.zLog.Fatal(args...)
}

func (d *DefaultLogger) Fatalf(template string, args ...interface{}) {
	d.zLog.Fatalf(template, args...)
}

func (d *DefaultLogger) Fatalw(msg string, keyAndValues ...interface{}) {
	d.zLog.Fatalw(msg, keyAndValues...)
}

var _ Logger = (*DefaultLogger)(nil)

func SetUp() error {
	cf := &LogConfig{}
	err := core.GetConfig("log", cf)
	if err != nil {
		return err
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	if cf.TimeFormat != "" {
		encoderCfg.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(t.Format(cf.TimeFormat))
		}
	}

	encoderCfg.EncodeLevel = cf.EncodeLevel.SwitchEncodeLevel()

	encoder := zapcore.NewJSONEncoder(encoderCfg)

	logLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= cf.Level.SwitchLevel()
	})

	writeSync := zapcore.NewMultiWriteSyncer(cf.Mode.SwitchWriter(cf)...)
	newCore := zapcore.NewCore(encoder, writeSync, logLevel)

	zl := zap.New(newCore, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()

	defaultLogger = &DefaultLogger{
		zLog: zl,
	}

	return nil

}

func Log() Logger {
	return defaultLogger
}

// SetLog 设置自定义Log
func SetLog(l Logger) {
	defaultLogger = l
}
