package logf

import (
	"fmt"
	"log"
	"os"
)

var (
	errorLog = log.New(os.Stdout, "\033[31m[ERROR]\033[0m", log.LstdFlags|log.Lshortfile)

	Error = func(v ...any) {
		if v != nil && v[0] != nil {
			_ = errorLog.Output(2, fmt.Sprintln(v...))
		}
	}
	Errorf = func(format string, v ...any) {
		if v != nil && v[0] != nil {
			_ = errorLog.Output(2, fmt.Sprintf(format, v...))
		}
	}

	Panic = func(v ...any) {
		if v != nil && v[0] != nil {
			s := fmt.Sprintln(v...)
			_ = errorLog.Output(2, s)
			panic(s)
		}
	}
	Panicf = func(format string, v ...any) {
		if v != nil && v[0] != nil {
			s := fmt.Sprintf(format, v...)
			_ = errorLog.Output(2, s)
			panic(s)
		}
	}

	Fatal = func(v ...any) {
		if v != nil && v[0] != nil {
			_ = errorLog.Output(2, fmt.Sprintln(v...))
			os.Exit(1)
		}
	}
	Fatalf = func(format string, v ...any) {
		if v != nil && v[0] != nil {
			_ = errorLog.Output(2, fmt.Sprintf(format, v...))
			os.Exit(1)
		}
	}
)

func init() {
	errorLog.SetOutput(os.Stdout)
}
