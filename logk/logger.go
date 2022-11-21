package logk

import (
	"github.com/cyj19/gowalk/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
	"time"
)

var (
	dLogger Logger
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
	Skip(int) Logger      // 返回溯源n层的Logger
	PoolGet() interface{} // 从对象池中获取实际的Logger对象
	PoolPut(interface{})  // 把实例对象放回对象池
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
	zLog *zap.Logger
}

func (d *defaultLogger) Debug(args ...interface{}) {
	d.zLog.Sugar().Debug(args...)
}

func (d *defaultLogger) Debugf(template string, args ...interface{}) {
	d.zLog.Sugar().Debugf(template, args...)
}

func (d *defaultLogger) Debugw(msg string, keyAndValues ...interface{}) {
	d.zLog.Sugar().Debugw(msg, keyAndValues...)
}

func (d *defaultLogger) Info(args ...interface{}) {
	d.zLog.Sugar().Info(args...)
}

func (d *defaultLogger) Infof(template string, args ...interface{}) {
	d.zLog.Sugar().Infof(template, args...)
}

func (d *defaultLogger) Infow(msg string, keyAndValues ...interface{}) {
	d.zLog.Sugar().Infow(msg, keyAndValues...)
}

func (d *defaultLogger) Warn(args ...interface{}) {
	d.zLog.Sugar().Warn(args...)
}

func (d *defaultLogger) Warnf(template string, args ...interface{}) {
	d.zLog.Sugar().Warnf(template, args...)
}

func (d *defaultLogger) Warnw(msg string, keyAndValues ...interface{}) {
	d.zLog.Sugar().Warnw(msg, keyAndValues...)
}

func (d *defaultLogger) Error(args ...interface{}) {
	d.zLog.Sugar().Error(args...)
}

func (d *defaultLogger) Errorf(template string, args ...interface{}) {
	d.zLog.Sugar().Errorf(template, args...)
}

func (d *defaultLogger) Errorw(msg string, keyAndValues ...interface{}) {
	d.zLog.Sugar().Errorw(msg, keyAndValues...)
}

func (d *defaultLogger) Panic(args ...interface{}) {
	d.zLog.Sugar().Panic(args...)
}

func (d *defaultLogger) Panicf(template string, args ...interface{}) {
	d.zLog.Sugar().Panicf(template, args...)
}

func (d *defaultLogger) Panicw(msg string, keyAndValues ...interface{}) {
	d.zLog.Sugar().Panicw(msg, keyAndValues...)
}

func (d *defaultLogger) Fatal(args ...interface{}) {
	d.zLog.Sugar().Fatal(args...)
}

func (d *defaultLogger) Fatalf(template string, args ...interface{}) {
	d.zLog.Sugar().Fatalf(template, args...)
}

func (d *defaultLogger) Fatalw(msg string, keyAndValues ...interface{}) {
	d.zLog.Sugar().Fatalw(msg, keyAndValues...)
}

func (d *defaultLogger) Skip(k int) Logger {
	// clone
	dc := d.PoolGet().(*defaultLogger)
	*dc = *d
	l := dc.zLog.WithOptions(zap.AddCallerSkip(k))
	dc.zLog = l
	return dc
}

func (d *defaultLogger) PoolGet() interface{} {
	return defaultLoggerPool.Get()
}

func (d *defaultLogger) PoolPut(v interface{}) {
	dc := v.(*defaultLogger)
	// 清空临时对象的字段
	dc.zLog = nil
	defaultLoggerPool.Put(v)
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

	l := zap.New(newCore, zap.AddCaller(), zap.AddCallerSkip(1))
	dLogger = &defaultLogger{
		zLog: l,
	}

	return nil

}

// GetLogger 获取日志对象
func GetLogger() Logger {
	logMu.Lock()
	defer logMu.Unlock()
	return dLogger
}

// SetLogger 设置自定义Logger，同时需要注意更新组件的Logger
func SetLogger(l Logger) {
	logMu.Lock()
	defer logMu.Unlock()
	dLogger = l
}

func Revert() {
	logMu.Lock()
	defer logMu.Unlock()
	logCfg := LogConfig{}
	_ = config.GetConfig("log", &logCfg)
	_ = SetupLog("./", logCfg)
}
