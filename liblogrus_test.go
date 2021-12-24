package liblogrus

import (
	"testing"

	"github.com/sgostarter/i/logger"
)

func Test1(t *testing.T) {
	rl := NewLogrus()
	rl.SetLevel(logger.LevelInfo)
	log := logger.NewWrapper(NewLogrus())

	log.WithFields(logger.FieldString("s1k", "s1v")).Info("ddd")
	log.WithFields(logger.FieldString("s1k", "s1v"), logger.FieldAny("int", 9887)).Info("ddd")
}
