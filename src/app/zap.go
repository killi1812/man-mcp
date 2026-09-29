package app

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func prodLoggerSetup() error {
	consoleLogLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.InfoLevel
	})

	// log output
	consoleLogFile := zapcore.Lock(os.Stdout)

	// log configuration no date time and location, just level
	consoleLogConfig := zap.NewProductionEncoderConfig()
	consoleLogConfig.EncodeTime = nil
	consoleLogConfig.EncodeCaller = nil
	// consoleLogConfig.EncodeLevel = nil
	consoleLogConfig.LevelKey = ""

	consoleLogEncoder := zapcore.NewConsoleEncoder(consoleLogConfig)

	/*
		// file log, text
		// log level
		fileLogLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return level >= zapcore.InfoLevel
		})

		logPath := filepath.Join(config.LOG_FOLDER, config.LOG_FILE)
		lumberjackLogger := lumberjack.Logger{
			Filename:   logPath,
			MaxSize:    config.LOG_FILE_MAX_SIZE,    // size in MB
			MaxAge:     config.LOG_FILE_MAX_AGE,     // maximum number of days to retain old log files
			MaxBackups: config.LOG_FILE_MAX_BACKUPS, // maximum number of old log files to retain
			LocalTime:  true,                        // time used for formatting the timestamps
			Compress:   false,
		}
		fileLogFile := zapcore.Lock(zapcore.AddSync(&lumberjackLogger))
		// log configuration
		fileLogConfig := zap.NewProductionEncoderConfig()
		// configure keys
		fileLogConfig.TimeKey = "timestamp"
		fileLogConfig.MessageKey = "message"
		// configure types
		fileLogConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		fileLogConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		// create encoder
		fileLogEncoder := zapcore.NewConsoleEncoder(fileLogConfig)
	*/
	// setup zap
	// duplicate log entries into multiple cores
	core := zapcore.NewTee(
		zapcore.NewCore(consoleLogEncoder, consoleLogFile, consoleLogLevel),
		//		zapcore.NewCore(fileLogEncoder, fileLogFile, fileLogLevel),
	)

	// create logger from core
	// options = annotate message with the filename, line number, and function name
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()

	// replace global logger
	_ = zap.ReplaceGlobals(logger)

	return nil
}

func devLoggerSetup() error {
	logger, err := zap.NewDevelopment(zap.AddStacktrace(zap.PanicLevel))
	if err != nil {
		return err
	}
	_ = zap.ReplaceGlobals(logger)

	return nil
}
