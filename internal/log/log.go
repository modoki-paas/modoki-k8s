package log

import (
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Logger struct {
	*logrus.Logger

	opts []grpc_logrus.Option
}

func New() *Logger {
	logger := logrus.New()

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
