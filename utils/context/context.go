package context

import (
	"context"

	"go.uber.org/zap"
)

type key int

const (
	keyRequestID key = iota + 1
	keyLogger
	keyUserID
)

// WithRequestID stores the request id in the context.
// It returns the context with the request id.
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, keyRequestID, requestID)
}

// RequestID returns the request id associated with the context.
// If there's no request id, it will return empty string.
func RequestID(ctx context.Context) string {
	v, _ := ctx.Value(keyRequestID).(string)
	return v
}

// WithLogger stores the zap logger in the context.
// It returns the context with the zap Logger.
func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, keyLogger, logger)
}

// Logger returns the zap logger associated with the context.
// If there's no zap logger, it will return a global Logger.
func Logger(ctx context.Context) *zap.Logger {
	if v, ok := ctx.Value(keyLogger).(*zap.Logger); ok {
		return v
	}
	return zap.L()
}

// WithUserID stores the user id in the context.
// It returns the context with the user id.
func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, keyUserID, userID)
}

// UserID returns the user id associated with the context.
// If there's no user id, it will return empty string.
func UserID(ctx context.Context) string {
	id, _ := ctx.Value(keyUserID).(string)
	return id
}
