package zapLogger

import (
	"github.com/qujing226/quantum_post/pkg/log"
	"go.uber.org/zap"
)

type ZapLogger struct {
	l *zap.Logger
}

func NewZapLogger(l *zap.Logger) log.Logger {
	return &ZapLogger{
		l: l,
	}
}
func (z ZapLogger) Debug(msg string, args ...log.Field) {
	z.l.Debug(msg, toZapFields(args)...)
}

func (z ZapLogger) Info(msg string, args ...log.Field) {
	z.l.Info(msg, toZapFields(args)...)
}

func (z ZapLogger) Warn(msg string, args ...log.Field) {
	z.l.Warn(msg, toZapFields(args)...)
}

func (z ZapLogger) Error(msg string, args ...log.Field) {
	z.l.Error(msg, toZapFields(args)...)
}

func toZapFields(args []log.Field) []zap.Field {
	res := make([]zap.Field, 0, len(args))
	for _, arg := range args {
		res = append(res, zap.Any(arg.Key, arg.Value))
	}
	return res
}
