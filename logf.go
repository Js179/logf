package logf

import (
	"io"
	"log"
	"os"
	"sync"
)

var (
	errorLog = log.New(os.Stdout, "\033[31m[ERROR]\033[0m", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[32m[INFO ]\033[0m", log.LstdFlags|log.Lshortfile)
	warnLog  = log.New(os.Stdout, "\033[33m[WARN ]\033[0m", log.LstdFlags|log.Lshortfile)
	debugLog = log.New(os.Stdout, "\033[34m[DEBUG]\033[0m", log.LstdFlags|log.Lshortfile)
	panicLog = log.New(os.Stdout, "\033[35m[PANIC]\033[0m", log.LstdFlags|log.Lshortfile)
	mu       sync.Mutex
)

var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
	Warn   = warnLog.Printf
	Debug  = debugLog.Printf
	Panic  = panicLog.Panicf
)

type Level uint8

const (
	ErrorLevel Level = iota
	WarnLevel
	InfoLevel
	DebugLevel
)

func init() {
	SetLevel(DebugLevel)
}

func SetLevel(level Level) {
	mu.Lock()
	defer mu.Unlock()

	panicLog.SetOutput(os.Stdout)
	errorLog.SetOutput(os.Stdout)
	warnLog.SetOutput(os.Stdout)
	debugLog.SetOutput(io.Discard)

	switch level {
	case DebugLevel:
		infoLog.SetOutput(os.Stdout)
		debugLog.SetOutput(os.Stdout)
	case InfoLevel:
		infoLog.SetOutput(os.Stdout)
	case ErrorLevel:
		warnLog.SetOutput(io.Discard)
		infoLog.SetOutput(io.Discard)
	case WarnLevel:
		infoLog.SetOutput(io.Discard)
	}
}
