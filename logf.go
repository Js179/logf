package logf

import (
	"io"
	"log"
	"os"
	"sync"
)

type Level int8

type option struct {
	Level
	io.Writer
	*sync.Mutex
	loggers []*log.Logger
}

var opt *option

const (
	ErrorLevel Level = iota
	WarnLevel
	InfoLevel
	DebugLevel
)

func init() {
	opt = &option{
		DebugLevel,
		os.Stdout,
		new(sync.Mutex),
		[]*log.Logger{errorLog, warnLog, infoLog, debugLog},
	}
}

func SetLevel(level Level) {
	set(level, opt.Writer)
}

// Stdout 终端打印日志
func Stdout() {
	set(opt.Level, os.Stdout)
}

// FileExport 日志文件输出
func FileExport(writer io.Writer) {
	set(opt.Level, writer)
}

// FileExportAndStdout 日志文件输出、终端打印
func FileExportAndStdout(writer io.Writer) {
	set(opt.Level, io.MultiWriter(os.Stdout, writer))
}

func set(level Level, writer io.Writer) {
	opt.Mutex.Lock()
	defer opt.Mutex.Unlock()
	opt.Level = level
	opt.Writer = writer

	for i := range opt.loggers {
		opt.loggers[i].SetOutput(getWriter(opt.Level, Level(i), opt.Writer))
	}
}

func getWriter(level, base Level, writer io.Writer) io.Writer {
	if level < base {
		return io.Discard
	}
	return writer
}
