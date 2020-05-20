package log

import (
	"context"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Fields is a weak typedef for logrus.Fields
type Fields = logrus.Fields

// Logger wraps logger for modoki
type Logger struct {
	*logrus.Logger

	opts []grpc_logrus.Option
}

// New initializes a new logger
func New() *Logger {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	return &Logger{
		Logger: logger,
		opts:   nil,
	}
}

// UnaryInterceptor returns an unary server interceptor for gRPC
func (l *Logger) UnaryInterceptor() grpc.UnaryServerInterceptor {
	entry := logrus.NewEntry(l.Logger)

	grpc_logrus.ReplaceGrpcLogger(entry)

	return grpc_logrus.UnaryServerInterceptor(entry, l.opts...)
}

// StreamInterceptor returns an unary server interceptor for gRPC
func (l *Logger) StreamInterceptor() grpc.StreamServerInterceptor {
	entry := logrus.NewEntry(l.Logger)

	grpc_logrus.ReplaceGrpcLogger(entry)

	return grpc_logrus.StreamServerInterceptor(entry, l.opts...)
}

// Extract gets logger for connection
func Extract(ctx context.Context) *logrus.Entry {
	return grpc_logrus.Extract(ctx)
}
