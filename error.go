package logf

import (
	"log"
	"os"
)

var (
	errorLog = log.New(os.Stdout, "\033[31m[ERROR]\033[0m", log.LstdFlags|log.Lshortfile)
	Error    = errorLog.Println
	Errorf   = errorLog.Printf

	Panic = func(v ...any) {
		if v != nil && v[0] != nil {
			errorLog.Panicln(v...)
		}
	}
	Panicf = func(format string, v ...any) {
		if v != nil && v[0] != nil {
			errorLog.Panicf(format, v...)
		}
	}

	Fatal = func(v ...any) {
		if v != nil && v[0] != nil {
			errorLog.Fatalln(v...)
		}
	}
	Fatalf = func(format string, v ...any) {
		if v != nil && v[0] != nil {
			errorLog.Fatalf(format, v...)
		}
	}
)

func init() {
	errorLog.SetOutput(os.Stdout)
}
