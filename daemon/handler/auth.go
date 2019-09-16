package handler

import (
	"context"
	"fmt"

	"strings"

	"github.com/modoki-paas/modoki-k8s/daemon/store"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Authorizer interface {
	Authorize(ctx context.Context, id int, route string) error
}

type AuthorizerInterceptor struct {
	db *store.DB
}

func (ai *AuthorizerInterceptor) getTokenFromHeader(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", xerrors.Errorf("metadata not found")
	}

	headers := md.Get("Authorization")

	if len(headers) == 0 {
		return "", xerrors.Errorf("metadata not found")
	}

	key := strings.TrimPrefix(headers[0], "Bearer ")

	return key, nil
}

func (ai *AuthorizerInterceptor) getIDFromToken(ctx context.Context) (int, error) {

}

// UnaryServerInterceptor はリクエストごとの認可を行う、unary サーバーインターセプターを返す。
func UnaryServerInterceptor(db *store.DB) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		var err error
		if srv, ok := info.Server.(Authorizer); ok {
			err = srv.Authorize(ctx, info.FullMethod)
		} else {
			return nil, fmt.Errorf("each service should implement an authorization")
		}
		if err != nil {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return handler(ctx, req)
	}
}
