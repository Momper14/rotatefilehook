# RotateFileHook

This is a simple hook for logrus to write log files using https://github.com/natefinch/lumberjack


```go
import (
  "github.com/ik5/rotatefilehook"
)

rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
    Filename: "logfile.log",
    MaxSize: 5,
    MaxBackups: 7,
    MaxAge: 7,
    Level: logrus.LevelDebug,
    Formatter: logrus.TextFormatter,
})
if err != nil {
  panic(err)
}

log.Hooks.Add(rotateFileHook)

err = rotateFileHook.Rotate() // To force rotation
if err != nil {
  panic(err)
}
```
