package logf

import (
	"io"
	"os"
	"sync"
)

var mu sync.Mutex

type Level uint8

const (
	ErrorLevel Level = iota
	WarnLevel
	InfoLevel
	DebugLevel
)

func SetLevel(level Level) {
	mu.Lock()
	defer mu.Unlock()

	warnLog.SetOutput(getLevel(level, WarnLevel))
	infoLog.SetOutput(getLevel(level, InfoLevel))
	debugLog.SetOutput(getLevel(level, DebugLevel))
}

func getLevel(level, base Level) io.Writer {
	if level < base {
		return io.Discard
	}
	return os.Stdout
}

func init() {
	SetLevel(DebugLevel)
}
