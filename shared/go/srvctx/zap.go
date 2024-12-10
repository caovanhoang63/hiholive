package srvctx

import (
	"context"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
	l.logger.Debug(fmt.Sprint(a...))
}

func (l *ZapLogger) Debugln(a ...any) {
	l.logger.Debug(fmt.Sprintln(a...))
}

func (l *ZapLogger) Debugf(s string, a ...any) {
	l.logger.Debug(fmt.Sprintf(s, a...))
}

func (l *ZapLogger) Info(a ...any) {
	l.logger.Info(fmt.Sprint(a...))
}

func (l *ZapLogger) Infoln(a ...any) {
	l.logger.Info(fmt.Sprintln(a...))
}

func (l *ZapLogger) Infof(s string, a ...any) {
	l.logger.Info(fmt.Sprintf(s, a...))
}

func (l *ZapLogger) Warn(a ...any) {
	l.logger.Warn(fmt.Sprint(a...))
}

func (l *ZapLogger) Warnln(a ...any) {
	l.logger.Warn(fmt.Sprintln(a...))
}

func (l *ZapLogger) Warnf(s string, a ...any) {
	l.logger.Warn(fmt.Sprintf(s, a...))
}

func (l *ZapLogger) Error(a ...any) {
	l.logger.Error(fmt.Sprint(a...))
}

func (l *ZapLogger) Errorln(a ...any) {
	l.logger.Error(fmt.Sprintln(a...))
}

func (l *ZapLogger) Errorf(s string, a ...any) {
	l.logger.Error(fmt.Sprintf(s, a...))
}

func (l *ZapLogger) Fatal(a ...any) {
	l.logger.Fatal(fmt.Sprint(a...))
}

func (l *ZapLogger) Fatalln(a ...any) {
	l.logger.Fatal(fmt.Sprintln(a...))
}

func (l *ZapLogger) Fatalf(s string, a ...any) {
	l.logger.Fatal(fmt.Sprintf(s, a...))
}

func (l *ZapLogger) Panic(a ...any) {
	l.logger.Panic(fmt.Sprint(a...))
}

func (l *ZapLogger) Panicln(a ...any) {
	l.logger.Panic(fmt.Sprintln(a...))
}

func (l *ZapLogger) Panicf(s string, a ...any) {
	l.logger.Panic(fmt.Sprintf(s, a...))
}

func (l *ZapLogger) WithField(field Field) Logger {
	newLogger := &ZapLogger{
		logger:   l.logger.With(zap.Any("arg", field)),
		ctx:      l.ctx,
		logLevel: l.logLevel,
	}
	return newLogger
}

func (l *ZapLogger) WithSrc(skip int) Logger {
	newLogger := &ZapLogger{
		logger:   l.logger.WithOptions(zap.AddCaller(), zap.AddCallerSkip(skip)),
		ctx:      l.ctx,
		logLevel: l.logLevel,
	}
	return newLogger
}

func (l *ZapLogger) GetLevel() string {
	return l.logLevel
}

func (l *ZapLogger) GetLogger(prefix string) Logger {
	return &ZapLogger{
		logger:   l.logger.With(zap.String("prefix", prefix)),
		ctx:      l.ctx,
		logLevel: l.logLevel,
	}
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

func (l *ZapLogger) Activate(serviceContext ServiceContext) error {
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
	logger := zap.New(core)
	return &ZapLogger{logger: logger, ctx: ctx}
}
