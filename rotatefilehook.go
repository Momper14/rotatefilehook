package rotatefilehook

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// RotateFileConfig holds basic information for using lumberjack and logrus hooks
type RotateFileConfig struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Level      logrus.Level
	Formatter  logrus.Formatter
}

// RotateFileHook holds hook information for logrus
type RotateFileHook struct {
	Config    RotateFileConfig
	LogWriter *lumberjack.Logger
}

// NewRotateFileHook initialize a new logrus.Hook or return an error
func NewRotateFileHook(config RotateFileConfig) (*RotateFileHook, error) {

	hook := &RotateFileHook{
		Config: config,
		LogWriter: &lumberjack.Logger{
			Filename:   config.Filename,
			MaxSize:    config.MaxSize,
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,
		},
	}

	return hook, nil
}

// Levels implements the Levels interface method of logrus Hook
func (hook *RotateFileHook) Levels() []logrus.Level {
	return logrus.AllLevels[:hook.Config.Level+1]
}

// Fire implements the Fire interface method of logrus Hook
func (hook *RotateFileHook) Fire(entry *logrus.Entry) (err error) {
	b, err := hook.Config.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = hook.LogWriter.Write(b)
	return err
}

// Rotate by request a log file (calling of SIGHUP for example)
func (hook *RotateFileHook) Rotate() error {
	return hook.LogWriter.Rotate()
}
