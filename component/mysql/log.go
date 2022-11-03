package mysql

import (
	"context"
	"errors"
	"fmt"
	"github.com/cyj19/gowalk/core/logx"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

// 实现gorm.Logger接口，copy源码，把标准库替换为自定义的Logger即可

type SqlLogger struct {
	gormlogger.Config
	xLog                                         logx.Logger
	debugStr, infoStr, warnStr, errStr, fatalStr string
	traceStr, traceErrStr, traceWarnStr          string
}

func New(xLog logx.Logger, cf gormlogger.Config) *SqlLogger {
	var (
		debugStr     = "%s\n[debug]"
		infoStr      = "%s[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		fatalStr     = "%s[fatal]"
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)

	if cf.Colorful {
		infoStr = gormlogger.Green + "%s\n" + gormlogger.Reset + gormlogger.Green + "[info] " + gormlogger.Reset
		warnStr = gormlogger.BlueBold + "%s\n" + gormlogger.Reset + gormlogger.Magenta + "[warn] " + gormlogger.Reset
		errStr = gormlogger.Magenta + "%s\n" + gormlogger.Reset + gormlogger.Red + "[error] " + gormlogger.Reset
		traceStr = gormlogger.Green + "%s\n" + gormlogger.Reset + gormlogger.Yellow + "[%.3fms] " + gormlogger.BlueBold + "[rows:%v]" + gormlogger.Reset + " %s"
		traceWarnStr = gormlogger.Green + "%s " + gormlogger.Yellow + "%s\n" + gormlogger.Reset + gormlogger.RedBold + "[%.3fms] " + gormlogger.Yellow + "[rows:%v]" + gormlogger.Magenta + " %s" + gormlogger.Reset
		traceErrStr = gormlogger.RedBold + "%s " + gormlogger.MagentaBold + "%s\n" + gormlogger.Reset + gormlogger.Yellow + "[%.3fms] " + gormlogger.BlueBold + "[rows:%v]" + gormlogger.Reset + " %s"
	}

	return &SqlLogger{
		xLog:         xLog,
		Config:       cf,
		debugStr:     debugStr,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		fatalStr:     fatalStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

func (l *SqlLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	nl := *l
	nl.LogLevel = level
	return &nl
}

func (l *SqlLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.xLog.Infof(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
}

func (l *SqlLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.xLog.Warnf(l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
}

func (l *SqlLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.xLog.Errorf(l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
}

func (l *SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= gormlogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gormlogger.Error && (!errors.Is(err, gormlogger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.xLog.Errorf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.xLog.Errorf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= gormlogger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.xLog.Warnf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.xLog.Warnf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.LogLevel == gormlogger.Info:
		sql, rows := fc()
		if rows == -1 {
			l.xLog.Infof(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.xLog.Infof(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
