package log

import (
	"context"

	"go.uber.org/zap"
)

type contextKey string

const loggerKey contextKey = "zap-logger"

// WithNewLogger builds logger and attach it to given context
func WithNewLogger(
	ctx context.Context,
	builders ...BuilderFunc,
) (context.Context, error) {
	b := NewBuilder()
	for _, builder := range builders {
		b = builder(b)
	}
	l, err := b.Build()
	if err != nil {
		return ctx, err
	}
	return WithLogger(ctx, l), nil
}

// WithNewLoggerForced does what [WithNewLogger] does but panics if fails
func WithNewLoggerForced(
	ctx context.Context,
	builders ...BuilderFunc,
) context.Context {
	b := NewBuilder()
	for _, builder := range builders {
		b = builder(b)
	}
	l := b.MustBuild()

	return WithLogger(ctx, l)
}

// WithLogger adds logger to context
func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

// FromContext extracts logger from context, or returns a nop logger if it was not found
func FromContext(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return logger
	}
	return zap.NewNop()
}
