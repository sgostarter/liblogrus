package liblogrus

import (
	"github.com/sgostarter/i/l"
	"github.com/sirupsen/logrus"
)

func NewLogrus() l.Logger {
	return NewLogrusEx(nil)
}

func NewLogrusEx(logger *logrus.Logger) l.Logger {
	if logger == nil {
		logger = logrus.New()
		logger.SetFormatter(new(logrus.JSONFormatter))
	}

	return &logrusImpl{
		rl: logrus.NewEntry(logger),
	}
}

type logrusImpl struct {
	rl *logrus.Entry
}

func (impl *logrusImpl) mapLevel(level l.Level) logrus.Level {
	switch level {
	case l.LevelFatal:
		return logrus.FatalLevel
	case l.LevelError:
		return logrus.ErrorLevel
	case l.LevelWarn:
		return logrus.WarnLevel
	case l.LevelInfo:
		return logrus.InfoLevel
	case l.LevelDebug:
		return logrus.DebugLevel
	}

	return logrus.FatalLevel
}

func (impl *logrusImpl) SetLevel(level l.Level) {
	impl.rl.Logger.SetLevel(impl.mapLevel(level))
}

func (impl *logrusImpl) WithFields(fields ...l.Field) l.Logger {
	fs := make(map[string]interface{})
	for _, field := range fields {
		fs[field.K] = field.V
	}

	return &logrusImpl{
		rl: impl.rl.WithFields(logrus.Fields(fs)),
	}
}

func (impl *logrusImpl) Log(level l.Level, a ...interface{}) {
	mLevel := impl.mapLevel(level)
	if mLevel == logrus.FatalLevel {
		impl.rl.Fatal(a...)
	} else {
		impl.rl.Log(mLevel, a...)
	}
}

func (impl *logrusImpl) Logf(level l.Level, format string, a ...interface{}) {
	mLevel := impl.mapLevel(level)
	if mLevel == logrus.FatalLevel {
		impl.rl.Fatalf(format, a...)
	} else {
		impl.rl.Logf(impl.mapLevel(level), format, a...)
	}
}
