package logf

import (
	"log"
	"os"
)

var (
	warnLog = log.New(os.Stdout, "\033[33m[WARN ]\033[0m", log.LstdFlags|log.Lshortfile)
	Warn    = warnLog.Println
	Warnf   = warnLog.Printf
)
