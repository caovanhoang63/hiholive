package zaplogger

import (
	"context"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"hiholive/shared/go/srvctx"
	"hiholive/shared/go/srvctx/components/loggerc"
	"log"
)

type ZapLogger struct {
	logger   *zap.Logger
	ctx      context.Context
	logLevel string
}

func (l *ZapLogger) Debug(a ...any) {
	if !l.logger.Level().Enabled(zapcore.DebugLevel) {
		return
	}
	fields := l.getContextCommonFields()
	l.logger.Debug(fmt.Sprint(a...), fields)
}

func (l *ZapLogger) Debugln(a ...any) {
	fields := l.getContextCommonFields()
	l.logger.Debug(fmt.Sprintln(a...), fields)
}

func (l *ZapLogger) Debugf(s string, a ...any) {
	fields := l.getContextCommonFields()
	l.logger.Debug(fmt.Sprintf(s, a...), fields)
}

func (l *ZapLogger) Info(a ...any) {
	fields := l.getContextCommonFields()
	l.logger.Info(fmt.Sprint(a...), fields)
}

func (l *ZapLogger) Infoln(a ...any) {
	fields := l.getContextCommonFields()
	l.logger.Info(fmt.Sprintln(a...), fields)
}

func (l *ZapLogger) Infof(s string, a ...any) {
	fields := l.getContextCommonFields()
	l.logger.Info(fmt.Sprintf(s, a...), fields)
}

func (l *ZapLogger) Warn(a ...any) {
	fields := l.getContextCommonFields()
	l.logger.Warn(fmt.Sprint(a...), fields)
}

func (l *ZapLogger) Warnln(a ...any) {
	fields := l.getContextCommonFields()
	l.logger.Warn(fmt.Sprintln(a...), fields)
}

func (l *ZapLogger) Warnf(s string, a ...any) {
	fields := l.getContextCommonFields()
	l.logger.Warn(fmt.Sprintf(s, a...), fields)
}

func (l *ZapLogger) Error(a ...any) {
	fields := l.getContextCommonFields()
	l.logger.Error(fmt.Sprint(a...), fields)
}

func (l *ZapLogger) Errorln(a ...any) {
	fields := l.getContextCommonFields()
	l.logger.Error(fmt.Sprintln(a...), fields)
}

func (l *ZapLogger) Errorf(s string, a ...any) {
	fields := l.getContextCommonFields()
	l.logger.Error(fmt.Sprintf(s, a...), fields)
}

func (l *ZapLogger) Fatal(a ...any) {
	fields := l.getContextCommonFields()
	l.logger.Fatal(fmt.Sprint(a...), fields)
}

func (l *ZapLogger) Fatalln(a ...any) {
	fields := l.getContextCommonFields()
	l.logger.Fatal(fmt.Sprintln(a...), fields)
}

func (l *ZapLogger) Fatalf(s string, a ...any) {
	fields := l.getContextCommonFields()
	l.logger.Fatal(fmt.Sprintf(s, a...), fields)
}

func (l *ZapLogger) Panic(a ...any) {
	fields := l.getContextCommonFields()
	l.logger.Panic(fmt.Sprint(a...), fields)
}

func (l *ZapLogger) Panicln(a ...any) {
	fields := l.getContextCommonFields()
	l.logger.Panic(fmt.Sprintln(a...), fields)
}

func (l *ZapLogger) Panicf(s string, a ...any) {
	fields := l.getContextCommonFields()
	l.logger.Panic(fmt.Sprintf(s, a...), fields)
}

func (l *ZapLogger) WithField(field loggerc.Field) loggerc.Logger {
	l.addContextCommonFields(field)
	return l
}

func (l *ZapLogger) WithSrc() loggerc.Logger {
	panic("implement me")
}

func (l *ZapLogger) GetLevel() string {
	return l.logLevel
}

func (l *ZapLogger) GetLogger(prefix string) loggerc.Logger {
	return DefaultLogger
}

func mustParseLevel(level string) zapcore.Level {
	lv, err := zapcore.ParseLevel(level)

	if err != nil {
		log.Fatal(err.Error())
	}

	return lv
}

func (l *ZapLogger) ID() string {
	return "loggerc"
}

func (l *ZapLogger) InitFlags() {
	level := l.logger.Level().String()
	flag.StringVar(&level, "log-level", "info", "Log level: panic | fatal | error | warn | info | debug | trace")
}

func (l *ZapLogger) Activate(serviceContext srvctx.ServiceContext) error {
	_ = mustParseLevel(l.logLevel)

	return nil
}

func (l *ZapLogger) Stop() error {
	return nil
}

var (
	DefaultLogger = NewZapLogger(context.Background(), zapcore.InfoLevel)
)

func NewZapLogger(ctx context.Context, level zapcore.LevelEnabler) *ZapLogger {
	encoder := getEncoderLog()
	sync := getWritterSync()
	core := zapcore.NewCore(encoder, sync, level)
	logger := zap.New(core, zap.AddCaller())
	return &ZapLogger{logger: logger, ctx: ctx}
}

func (l *ZapLogger) getContextCommonFields() zap.Field {
	if l.ctx != nil {
		if value := l.ctx.Value("commonFields"); value != nil {
			if fields, ok := value.(loggerc.Field); ok {
				return zap.Any("commonFields", fields)
			}
		}
	}
	return zap.Skip()
}

func (l *ZapLogger) addContextCommonFields(fields loggerc.Field) {
	if l.ctx != nil {
		for k, v := range l.ctx.Value("commonFields").(loggerc.Field) {
			if _, ok := fields[k]; !ok {
				fields[k] = v
			}
		}
	}
}
