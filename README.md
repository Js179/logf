# logf
golang log util of os.stdout

# How to get
```
go get github.com/js179/logf
```

# Usage
```Go
import "github.com/js179/logf"

func main() {
	// set log level
	// ErrorLevel
	// WarnLevel
	// InfoLevel
	// DebugLevel
	// logf.SetLevel(logf.InfoLevel)
	logf.Info("info...")
}
```
