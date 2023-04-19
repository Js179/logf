package logf

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func logs() {
	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")
	Panic(nil)
	Fatal(nil, nil)
}

func Test_Level(t *testing.T) {
	SetLevel(WarnLevel)
	logs()
}

func Test_Output(t *testing.T) {
	now := time.Now()
	date := fmt.Sprintf("%04d_%02d_%02d", now.Year(), now.Month(), now.Day())
	filename := fmt.Sprintf("logs%s%s%s", string(os.PathSeparator), date, ".log")
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 766)
	Panic(err)
	{
		FileExport(file)
		logs()
	}
	{
		FileExportAndStdout(file)
		logs()
	}
	{
		Stdout()
		logs()
	}
}
