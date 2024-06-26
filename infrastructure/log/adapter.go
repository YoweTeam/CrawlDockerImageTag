package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type zapAdapter struct {
	Path        string
	Level       string
	CataLog     string
	MaxFileSize int // unit(MB)
	MaxBackups  int
	MaxAge      int // save log, unit(day)
	Compress    bool
	Caller      bool // display caller

	logger *zap.Logger
	sugar  *zap.SugaredLogger
}

func (adapter *zapAdapter) copyWithCatalog(catalog string) *zapAdapter {
	a := *adapter
	a.CataLog = catalog
	a.sugar = a.logger.Sugar().Named(catalog)
	return &a
}

func (adapter *zapAdapter) setMaxFileSize(size int) {
	adapter.MaxFileSize = size
}

func (adapter *zapAdapter) setMaxBackups(n int) {
	adapter.MaxBackups = n
}

func (adapter *zapAdapter) setMaxAge(age int) {
	adapter.MaxAge = age
}

func (adapter *zapAdapter) setCompress(compress bool) {
	adapter.Compress = compress
}

func (adapter *zapAdapter) setCaller(caller bool) {
	adapter.Caller = caller
}

func NewZapAdapter(path string, level string, catalog string, maxFileSize, maxBackups, maxAge int) *zapAdapter {
	return &zapAdapter{
		Path:        path,
		Level:       level,
		CataLog:     catalog,
		MaxFileSize: maxFileSize,
		MaxBackups:  maxBackups,
		MaxAge:      maxAge,
		Compress:    true,
		Caller:      true,
	}
}

// createLumberjackHook Create LumberjackHook to cut and compress log files
func (adapter *zapAdapter) createLumberjackHook() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   adapter.Path,
		MaxSize:    adapter.MaxFileSize,
		MaxBackups: adapter.MaxBackups,
		MaxAge:     adapter.MaxAge,
		Compress:   adapter.Compress,
	}
}

// ParseLevel 解析日志级别
func ParseLevel(lv string) (zapcore.Level, error) {
	var level zapcore.Level
	switch lv {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		return 0, fmt.Errorf("parse log level failed for :%s", lv)
	}
	return level, nil
}

func (adapter *zapAdapter) Build() {
	w := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(adapter.createLumberjackHook()))

	level, err := ParseLevel(adapter.Level)
	// default log level
	if err != nil {
		level = zapcore.InfoLevel
	}

	conf := zap.NewProductionEncoderConfig()
	conf.EncodeTime = zapcore.ISO8601TimeEncoder
	cnf := zapcore.NewJSONEncoder(conf)
	core := zapcore.NewCore(cnf, w, level)

	adapter.logger = zap.New(core)
	if adapter.Caller {
		adapter.logger = adapter.logger.WithOptions(zap.AddCaller(), zap.AddCallerSkip(2))
	}
	adapter.sugar = adapter.logger.Sugar().Named(adapter.CataLog)
}

func (adapter *zapAdapter) SetCataLog(catalog string) *zap.SugaredLogger {
	return adapter.sugar.Named(catalog)
}

func (adapter *zapAdapter) WithFields(args ...interface{}) *zap.SugaredLogger {
	return adapter.sugar.With(args...)
}

func (adapter *zapAdapter) Debug(args ...interface{}) {
	adapter.sugar.Debug(args...)
}

func (adapter *zapAdapter) Info(args ...interface{}) {
	adapter.sugar.Info(args...)
}

func (adapter *zapAdapter) Warn(args ...interface{}) {
	adapter.sugar.Warn(args...)
}

func (adapter *zapAdapter) Error(args ...interface{}) {
	adapter.sugar.Error(args...)
}

func (adapter *zapAdapter) Panic(args ...interface{}) {
	adapter.sugar.Panic(args...)
}

func (adapter *zapAdapter) Fatal(args ...interface{}) {
	adapter.sugar.Fatal(args...)
}

func (adapter *zapAdapter) Debugf(template string, args ...interface{}) {
	adapter.sugar.Debugf(template, args...)
}

func (adapter *zapAdapter) Infof(template string, args ...interface{}) {
	adapter.sugar.Infof(template, args...)
}

func (adapter *zapAdapter) Warnf(template string, args ...interface{}) {
	adapter.sugar.Warnf(template, args...)
}

func (adapter *zapAdapter) Errorf(template string, args ...interface{}) {
	adapter.sugar.Errorf(template, args...)
}

func (adapter *zapAdapter) Panicf(template string, args ...interface{}) {
	adapter.sugar.Panicf(template, args...)
}

func (adapter *zapAdapter) Fatalf(template string, args ...interface{}) {
	adapter.sugar.Fatalf(template, args...)
}

func (adapter *zapAdapter) Debugw(msg string, keysAndValues ...interface{}) {
	adapter.sugar.Debugw(msg, keysAndValues...)
}

func (adapter *zapAdapter) Infow(msg string, keysAndValues ...interface{}) {
	adapter.sugar.Infow(msg, keysAndValues...)
}

func (adapter *zapAdapter) Warnw(msg string, keysAndValues ...interface{}) {
	adapter.sugar.Warnw(msg, keysAndValues...)
}

func (adapter *zapAdapter) Errorw(msg string, keysAndValues ...interface{}) {
	adapter.sugar.Errorw(msg, keysAndValues...)
}

func (adapter *zapAdapter) Panicw(msg string, keysAndValues ...interface{}) {
	adapter.sugar.Panicw(msg, keysAndValues...)
}

func (adapter *zapAdapter) Fatalw(msg string, keysAndValues ...interface{}) {
	adapter.sugar.Fatalw(msg, keysAndValues...)
}
