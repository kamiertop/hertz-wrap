package log

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/json"
	zlog "github.com/hertz-contrib/logger/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"hertz/pkg/config"
	"hertz/pkg/consts"
)

var Logger = new(zap.Logger)

// InitLogger Set default logger and global zap logger
func InitLogger() {
	var cores []zlog.CoreConfig
	if config.Cfg().System.Env == consts.DevelopmentMode {
		cores = append(cores, zlog.CoreConfig{
			Enc: zapcore.NewConsoleEncoder(consoleEncoder()),
			Ws:  zapcore.AddSync(stdWriter()),
			Lvl: level(),
		})
	} else {
		cores = append(cores, zlog.CoreConfig{
			Enc: zapcore.NewConsoleEncoder(jsonEncoder()),
			Ws:  zapcore.AddSync(fileWriter()),
			Lvl: level(),
		})
	}
	logger := zlog.NewLogger(zlog.WithCores(cores...), zlog.WithZapOptions(zap.AddCaller(), zap.AddCallerSkip(1)))

	Logger = logger.Logger()

	hlog.SetLogger(logger)
}

func fileWriter() io.Writer {
	return &lumberjack.Logger{
		Filename:   config.Cfg().Log.FileName,
		MaxSize:    20,   // A file can be up to 20M.
		MaxBackups: 5,    // Save up to 5 files at the same time.
		MaxAge:     10,   // A file can exist for a maximum of 10 days.
		Compress:   true, // Compress with gzip.
	}
}

func stdWriter() io.Writer {
	return os.Stdout
}

func level() zap.AtomicLevel {
	l := zap.NewAtomicLevel() // default level is InfoLevel
	switch strings.ToLower(config.Cfg().Log.Level) {
	case "debug":
		l.SetLevel(zapcore.DebugLevel)
	case "info":
		l.SetLevel(zapcore.InfoLevel)
	case "warn":
		l.SetLevel(zapcore.WarnLevel)
	case "error":
		l.SetLevel(zapcore.ErrorLevel)
	case "panic":
		l.SetLevel(zapcore.PanicLevel)
	}

	return l
}

func consoleEncoder() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey: "msg",
		LevelKey:   "level",
		TimeKey:    "ts",
		NameKey:    "name",
		CallerKey:  "caller",
		//FunctionKey:    "fn",
		StacktraceKey:  "st",
		SkipLineEnding: false,
		LineEnding:     "",
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout(time.DateTime),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		NewReflectedEncoder: func(writer io.Writer) zapcore.ReflectedEncoder {
			return json.NewEncoder(writer)
		},
		ConsoleSeparator: " ",
	}
}

func jsonEncoder() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "ts",
		NameKey:        "name",
		CallerKey:      "caller",
		FunctionKey:    "fn",
		StacktraceKey:  "st",
		SkipLineEnding: false,
		LineEnding:     "",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout(time.DateTime),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
		NewReflectedEncoder: func(writer io.Writer) zapcore.ReflectedEncoder {
			return json.NewEncoder(writer)
		},
	}
}

func Debug(msg string, fields ...zapcore.Field) {
	Logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zapcore.Field) {
	Logger.Info(msg, fields...)
}

func Infof(format string, v ...any) {
	Logger.Sugar().Infof(format, v...)
}

func Warn(msg string, fields ...zapcore.Field) {
	Logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	Logger.Error(msg, fields...)
}

func Panic(msg string, fields ...zapcore.Field) {
	Logger.Panic(msg, fields...)
}

func Sync() error {
	return Logger.Sync()
}

func Sugar() *zap.SugaredLogger {
	return Logger.Sugar()
}
