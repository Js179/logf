package logf

import (
	"log"
	"os"
)

var (
	debugLog = log.New(os.Stdout, "\033[34m[DEBUG]\033[0m", log.LstdFlags|log.Lshortfile)
	Debug    = debugLog.Println
	Debugf   = debugLog.Printf
)
