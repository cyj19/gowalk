package logk

import "sync"

var defaultLoggerPool = sync.Pool{
	New: func() interface{} {
		return new(defaultLogger)
	},
}
