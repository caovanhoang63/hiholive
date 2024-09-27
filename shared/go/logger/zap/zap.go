package zaplogger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"hiholive/shared/go/logger"
)

type ZapLogger struct {
	logger *zap.Logger
	ctx    context.Context
}

func NewZapLogger(ctx context.Context, level zapcore.LevelEnabler) *ZapLogger {
	encoder := getEncoderLog()
	sync := getWritterSync()
	core := zapcore.NewCore(encoder, sync, level)
	logger := zap.New(core, zap.AddCaller())
	return &ZapLogger{logger: logger, ctx: ctx}
}

func (l *ZapLogger) DebugWithFields(msg string, fields logger.Field) {

	l.logger.Debug(msg, zap.Any("args", fields))
}

func (l *ZapLogger) InfoWithFields(msg string, fields logger.Field) {
	l.addContextCommonFields(fields)

	l.logger.Info(msg, zap.Any("args", fields))
}

func (l *ZapLogger) WarnWithFields(msg string, fields logger.Field) {
	l.addContextCommonFields(fields)
	l.logger.Warn(msg, zap.Any("args", fields))
}

func (l *ZapLogger) ErrorWithFields(msg string, fields logger.Field) {
	l.addContextCommonFields(fields)

	l.logger.Error(msg, zap.Any("args", fields))
}

func (l *ZapLogger) FatalWithFields(msg string, fields logger.Field) {
	l.addContextCommonFields(fields)

	l.logger.Fatal(msg, zap.Any("args", fields))
}

func (l *ZapLogger) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l *ZapLogger) Info(msg string) {
	l.logger.Info(msg)

}

func (l *ZapLogger) Warn(msg string) {
	l.logger.Warn(msg)

}

func (l *ZapLogger) Error(msg string) {
	l.logger.Error(msg)

}

func (l *ZapLogger) Fatal(msg string) {
	l.logger.Fatal(msg)
}

func (l *ZapLogger) addContextCommonFields(fields logger.Field) {
	if l.ctx != nil {
		for k, v := range l.ctx.Value("commonFields").(logger.Field) {
			if _, ok := fields[k]; !ok {
				fields[k] = v
			}
		}
	}
}
