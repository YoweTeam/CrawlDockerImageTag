package log

import (
	"github.com/YoweTeam/CrawlDockerImageTag/infrastructure/config"
	"github.com/YoweTeam/CrawlDockerImageTag/infrastructure/context"
	"testing"
)

func Test1(t *testing.T) {
	var ctx = context.NewBackgroundContext()
	Init(&LogSettings{
		Level:    config.GetStringOrDefault("log.level", DefaultLevel),
		Path:     config.GetStringOrDefault("log.path", DefaultPath),
		FileName: config.GetStringOrDefault("log.filename", DefaultFileName),
		CataLog:  config.GetStringOrDefault("log.catalog", DefaultCataLog),
		Caller:   config.GetBoolOrDefault("log.caller", DefaultCaller),
	})
	logger.WithContext(ctx).Info("hello world")
	WithContext(ctx).Info("hello world 233333")
	Sync()
}

func Test2(t *testing.T) {
	var ctx = context.NewBackgroundContext()
	Init(&LogSettings{
		Level:    config.GetStringOrDefault("log.level", DefaultLevel),
		Path:     config.GetStringOrDefault("log.path", DefaultPath),
		FileName: config.GetStringOrDefault("log.filename", DefaultFileName),
		CataLog:  config.GetStringOrDefault("log.catalog", DefaultCataLog),
		Caller:   config.GetBoolOrDefault("log.caller", DefaultCaller),
	})

	WithContext(ctx).Debug("test Debug")
	WithContext(ctx).Debugf("test Debugf")
	WithContext(ctx).Debugw("test", "Debugw")

	WithContext(ctx).Info("test Info")
	WithContext(ctx).Infof("test Infof")
	WithContext(ctx).Infow("test", "Infow")

	WithContext(ctx).Warn("test Warn")
	WithContext(ctx).Warnf("test Warnf")
	WithContext(ctx).Warnw("test", "Warnw")

	WithContext(ctx).Error("test Error")
	WithContext(ctx).Errorf("test Errorf")
	WithContext(ctx).Errorw("test", "Errorw")

	Sync()
}

func TestBenchmark(b *testing.B) {
	var ctx = context.NewBackgroundContext()
	b.ReportAllocs()
	Init(&LogSettings{
		Level:    config.GetStringOrDefault("log.level", DefaultLevel),
		Path:     config.GetStringOrDefault("log.path", DefaultPath),
		FileName: config.GetStringOrDefault("log.filename", DefaultFileName),
		CataLog:  config.GetStringOrDefault("log.catalog", DefaultCataLog),
		Caller:   config.GetBoolOrDefault("log.caller", DefaultCaller),
	})
	for i := 0; i < b.N; i++ {
		WithContext(ctx).Info(`{"level":"debug","ts":"2024-06-26T19:58:19.342+0800","logger":"serve","msg":"","res_body":"","referer":"","client_ip":"::1","method":"GET","req":"","res_headers":"","user_agent":"","st
ack":"","host":" NODOES-PUERSA","path":"/api/_sys/version","code":200,"latency":0}`)
	}
}
