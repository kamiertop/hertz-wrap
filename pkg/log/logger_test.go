package log

import (
	"testing"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"go.uber.org/zap"

	"hertz/pkg/config"
)

func TestInitLogger(t *testing.T) {
	_ = config.Init()
	Init()
	hlog.Info("info level")
	Info("info level")
	Info("info level with field", zap.String("string type", "string value"))
}
