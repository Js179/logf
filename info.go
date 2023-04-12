package logf

import (
	"log"
	"os"
)

var (
	infoLog = log.New(os.Stdout, "\033[32m[INFO ]\033[0m", log.LstdFlags|log.Lshortfile)
	Info    = infoLog.Println
	Infof   = infoLog.Printf
)
