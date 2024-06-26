package log

import (
	"fmt"
	contextBase "github.com/YoweTeam/CrawlDockerImageTag/infrastructure/context"

	"go.uber.org/zap"
)

type ILogger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	WithContext(context contextBase.Context) ILogger
	WithFields(args map[string]interface{}) ILogger
}

const (
	DefaultLevel       = "info"
	DefaultPath        = "./logs"
	DefaultFileName    = "app.log"
	DefaultCataLog     = "app"
	DefaultCaller      = false
	DefaultMaxFileSize = 1024 // 日志文件大小限制默认值，单位MB，超过该大小会自动压缩归档
	DefaultMaxBackups  = 3    // 日志归档文件个数默认值，超过后会滚动清理，仅当MaxBackups与MaxAge均为0时不再自动清理日志文件
	DefaultMaxAge      = 7    // 日志归档文件保留天数默认值，超过后会滚动清理，仅当MaxBackups与MaxAge均为0时不再自动清理日志文件
)

type LogSettings struct {
	Path        string
	Level       string
	CataLog     string
	FileName    string
	Caller      bool
	MaxFileSize int // 日志文件大小的最大值，单位(M)
	MaxBackups  int // 最多保留备份数
	MaxAge      int // 日志文件保存的时间，单位(天)

	adapter *zapAdapter
}
type xLogger struct {
	*zap.SugaredLogger
}

func (log *xLogger) WithFields(args map[string]interface{}) ILogger {
	var fields = make([]interface{}, 0)
	for k, v := range args {
		fields = append(fields, k)
		fields = append(fields, v)
	}
	log.SugaredLogger = log.SugaredLogger.With(fields...)
	return log
}

func (log *xLogger) WithContext(context contextBase.Context) ILogger {
	args := contextBase.ToMap(context)
	var fields = make([]interface{}, 0)
	for k, v := range args {
		fields = append(fields, k)
		fields = append(fields, v)
	}
	log.SugaredLogger = log.SugaredLogger.With(fields...)
	return log
}

var logger *LogSettings

// Init init logger
func Init(settings *LogSettings) error {
	logger = settings
	// instance
	logger.adapter = NewZapAdapter(fmt.Sprintf("%s/%s", logger.Path, logger.FileName), logger.Level, logger.CataLog,
		settings.MaxFileSize, settings.MaxBackups, settings.MaxAge)
	logger.adapter.setCaller(logger.Caller)
	logger.adapter.Build()

	return nil
}

// Sync flushes buffer, if any
func Sync() {
	if logger == nil {
		return
	}

	logger.adapter.logger.Sync()
}

// WithCataLog get log object with catalog
func WithCataLog(catalog string) *LogSettings {
	if logger == nil {
		return nil
	}
	innerLogger := *logger
	innerLogger.CataLog = catalog
	innerLogger.adapter = logger.adapter.copyWithCatalog(catalog)
	return &innerLogger
}

func (settings *LogSettings) With(args ...interface{}) ILogger {
	return &xLogger{SugaredLogger: settings.adapter.WithFields(args...)}
}

func (settings *LogSettings) WithContext(ctx contextBase.Context) ILogger {
	var fields = getFields(ctx)
	return &xLogger{SugaredLogger: settings.adapter.WithFields(fields...)}
}

func (settings *LogSettings) WithFields(args map[string]interface{}) ILogger {

	var fields = make([]interface{}, 0)
	for k, v := range args {
		fields = append(fields, k)
		fields = append(fields, v)
	}
	return &xLogger{SugaredLogger: settings.adapter.WithFields(fields...)}
}

func getFields(ctx contextBase.Context) []interface{} {
	args := contextBase.ToMap(ctx)
	var fields = make([]interface{}, 0)
	for k, v := range args {
		fields = append(fields, k)
		fields = append(fields, v)
	}
	return fields
}

func WithContext(context contextBase.Context) ILogger {
	var fields = getFields(context)
	return &xLogger{SugaredLogger: logger.adapter.WithFields(fields...)}
}

// Debug 使用方法：log.Debug("test")
func Debug(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Debug(args...)
}

// Debugf using: log.Debugf("test:%s", err)
func Debugf(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Debugf(template, args...)
}

// Debugw using: log.Debugw("test", "field1", "value1", "field2", "value2")
func Debugw(msg string, keysAndValues ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Debugw(msg, keysAndValues...)
}

func Info(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Info(args...)
}

func Infof(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Infow(msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Warnw(msg, keysAndValues...)
}

func Error(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Errorw(msg, keysAndValues...)
}

func Panic(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Panicw(msg, keysAndValues...)
}

func Fatal(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Fatalf(template, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Fatalw(msg, keysAndValues...)
}
