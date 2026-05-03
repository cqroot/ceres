package logger

import (
	"flag"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.SugaredLogger

func init() {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	newLog, err := cfg.Build()
	if err != nil {
		newLog = zap.NewNop()
	}
	log = newLog.Sugar()
}

var verboseFlag bool

func RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(&verboseFlag, "verbose", false, "enable verbose logging")
}

func sync() {
	_ = log.Sync()
}

func Debug(args ...interface{}) {
	if verboseFlag {
		log.Debug(args...)
	}
}

func Debugf(template string, args ...interface{}) {
	if verboseFlag {
		log.Debugf(template, args...)
	}
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(template string, args ...interface{}) {
	log.Infof(template, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	log.Warnf(template, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	log.Errorf(template, args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	log.Fatalf(template, args...)
}

func Sync() {
	sync()
}

func SetOutputPath(path string) error {
	if path == "" || path == "stdout" {
		return nil
	}

	if path == "stderr" {
		return nil
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	_ = file.Close()

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{path}
	cfg.ErrorOutputPaths = []string{path}
	newLog, err := cfg.Build()
	if err != nil {
		return err
	}
	log = newLog.Sugar()
	return nil
}
