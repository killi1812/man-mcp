package app

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func setupDevLogger() error {
	core := makeDevCore()
	return replaceGlobalWithCore(core)
}

func setupProdLogger() error {
	core := makeProdCore()
	return replaceGlobalWithCore(core)
}

func replaceGlobalWithCore(core zapcore.Core) error {
	logger := zap.New(core, zap.AddCaller())
	_ = zap.ReplaceGlobals(logger)
	return nil
}

func makeDevCore() zapcore.Core {
	level := zap.DebugLevel
	return buildTeeCore(zapcore.NewConsoleEncoder(devEncoderConfig()), level)
}

func makeProdCore() zapcore.Core {
	level := zap.InfoLevel
	return buildTeeCore(zapcore.NewJSONEncoder(prodEncoderConfig()), level)
}

func devEncoderConfig() zapcore.EncoderConfig {
	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return cfg
}

func prodEncoderConfig() zapcore.EncoderConfig {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	return cfg
}

func buildTeeCore(enc zapcore.Encoder, lvl zapcore.Level) zapcore.Core {
	stderrCore := zapcore.NewCore(enc, zapcore.Lock(os.Stderr), lvl)
	if LogFilePath == "" {
		return stderrCore
	}
	return zapcore.NewTee(stderrCore, makeFileCore(lvl))
}

func makeFileCore(lvl zapcore.Level) zapcore.Core {
	file, err := os.OpenFile(LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return zapcore.NewNopCore()
	}
	return zapcore.NewCore(zapcore.NewJSONEncoder(prodEncoderConfig()), zapcore.AddSync(file), lvl)
}
