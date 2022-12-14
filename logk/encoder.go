package logk

import "go.uber.org/zap/zapcore"

const (
	Capital EncodeLevel = iota
	CapitalColor
	Lowercase
	LowercaseColor
)

type EncodeLevel int8

func (e EncodeLevel) switchEncodeLevel() zapcore.LevelEncoder {
	switch e {
	case Capital:
		return zapcore.CapitalLevelEncoder
	case CapitalColor:
		return zapcore.CapitalColorLevelEncoder
	case Lowercase:
		return zapcore.LowercaseLevelEncoder
	case LowercaseColor:
		return zapcore.LowercaseColorLevelEncoder
	}

	return zapcore.CapitalLevelEncoder
}
