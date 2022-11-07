package gowalk

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
)

const (
	Console Mode = iota
	File
	ConsoleAndFile
)

type Mode int8

func (m Mode) SwitchWriter(cf *LogConfig) []zapcore.WriteSyncer {
	ws := make([]zapcore.WriteSyncer, 0)
	switch m {
	case Console:
		ws = append(ws, zapcore.AddSync(os.Stdout))
		return ws
	case File:
		ws = append(ws, zapcore.AddSync(GenLumberjackWriter(cf)))
		return ws
	case ConsoleAndFile:
		ws = append(ws, zapcore.AddSync(os.Stdout), zapcore.AddSync(GenLumberjackWriter(cf)))
		return ws
	}

	ws = append(ws, zapcore.AddSync(os.Stdout))

	return ws
}

func GenLumberjackWriter(cf *LogConfig) io.Writer {
	if cf.Name != "" {
		logName = cf.Name
	}
	filename := filepath.Join(WorkDir, "log", logName+".log")
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    cf.MaxSize,
		MaxAge:     cf.MaxAge,
		MaxBackups: cf.MaxBackups,
		Compress:   cf.Compress,
	}
}
