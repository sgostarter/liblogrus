package liblogrus

import (
	"testing"

	"github.com/sgostarter/i/l"
	"github.com/sgostarter/i/logger"
)

func Test1(t *testing.T) {
	rl := NewLogrus()
	rl.SetLevel(logger.LevelInfo)

	log := l.NewWrapper(NewLogrus())

	log.WithFields(l.StringField("s1k", "s1v")).Info("ddd")
	log.WithFields(l.StringField("s1k", "s1v"), l.IntField("int", 9887)).Info("ddd")
}
