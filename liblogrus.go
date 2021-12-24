package liblogrus

import (
	"github.com/sgostarter/i/logger"
	"github.com/sirupsen/logrus"
)

func NewLogrus() logger.Logger {
	return &logrusImpl{
		rl: logrus.NewEntry(logrus.New()),
	}
}

type logrusImpl struct {
	rl *logrus.Entry
}

func (impl *logrusImpl) mapLevel(level logger.Level) logrus.Level {
	switch level {
	case logger.LevelFatal:
		return logrus.FatalLevel
	case logger.LevelError:
		return logrus.ErrorLevel
	case logger.LevelWarn:
		return logrus.WarnLevel
	case logger.LevelInfo:
		return logrus.InfoLevel
	case logger.LevelDebug:
		return logrus.DebugLevel
	}

	return logrus.FatalLevel
}

func (impl *logrusImpl) SetLevel(level logger.Level) {
	impl.rl.Logger.SetLevel(impl.mapLevel(level))
}

func (impl *logrusImpl) WithFields(fields ...logger.Field) logger.Logger {
	fs := make(map[string]interface{})
	for _, field := range fields {
		fs[field.K] = field.V
	}

	return &logrusImpl{
		rl: impl.rl.WithFields(logrus.Fields(fs)),
	}
}

func (impl *logrusImpl) Log(level logger.Level, a ...interface{}) {
	impl.rl.Log(impl.mapLevel(level), a...)
}

func (impl *logrusImpl) Logf(level logger.Level, format string, a ...interface{}) {
	impl.rl.Logf(impl.mapLevel(level), format, a...)
}
