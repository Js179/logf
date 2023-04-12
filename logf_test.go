package logf

import "testing"

func logs() {
	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")
	Panic(nil)
	Fatal(nil, nil)
}

func Test_Warn(t *testing.T) {
	SetLevel(WarnLevel)
	logs()
}

func Test_Info(t *testing.T) {
	SetLevel(InfoLevel)
	logs()
}

func Test_Debug(t *testing.T) {
	SetLevel(DebugLevel)
	logs()
}
