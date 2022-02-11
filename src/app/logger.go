package app

import (
	"os"
	"path/filepath"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	debugLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.InfoLevel
	})
	infoLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel && lvl > zapcore.DebugLevel
	})
	warnLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})
	devEncoder = zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		TimeKey:     "ts",
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000000-07:00"))
		},
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
	prodEncoder = zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey: "msg",
		LevelKey:   "level",
		TimeKey:    "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000000-07:00"))
		},
	})
)

func NewLogger(run_mode, path string) *zap.Logger {
	if run_mode == "dev" {
		return getDevLogger()
	} else if run_mode == "debug" {
		return getDebugLogger()
	} else if run_mode == "prod" {
		return getProdLogger(path)
	}
	return getDevLogger()
}

func getDevLogger() *zap.Logger {
	var core = zapcore.NewTee(
		zapcore.NewCore(devEncoder, os.Stdout, zap.DebugLevel),
	)
	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func getDebugLogger() *zap.Logger {
	var core = zapcore.NewTee(
		zapcore.NewCore(devEncoder, os.Stdout, zap.DebugLevel),
	)
	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func getProdLogger(path string) *zap.Logger {
	var core = zapcore.NewTee(
		zapcore.NewCore(prodEncoder, getHook(path, "seed-info.log"), infoLevel),
		zapcore.NewCore(prodEncoder, getHook(path, "seed_err.log"), warnLevel),
	)
	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func getHook(path, filename string) zapcore.WriteSyncer {
	hook := &lumberjack.Logger{
		Filename:   filepath.Join(path, filename),
		MaxSize:    1,  // megabytes
		MaxBackups: 10, // megabytes
	}
	return zapcore.AddSync(hook)
}
