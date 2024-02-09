package log

import (
	"io"
	"os"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	zlog "github.com/hertz-contrib/logger/zap"
	"gopkg.in/natefinch/lumberjack.v2"

	"hertz/pkg/config"
)

// InitLogger Set default logger and global zap logger
func InitLogger() {
	logger := zlog.NewLogger()
	hlog.SetLogger(logger)
}

func fileWriter() io.Writer {
	return &lumberjack.Logger{
		Filename:   config.Conf.Log.FileName,
		MaxSize:    20,   // A file can be up to 20M.
		MaxBackups: 5,    // Save up to 5 files at the same time.
		MaxAge:     10,   // A file can exist for a maximum of 10 days.
		Compress:   true, // Compress with gzip.
	}
}

func stdWriter() io.Writer {
	return os.Stdout
}
