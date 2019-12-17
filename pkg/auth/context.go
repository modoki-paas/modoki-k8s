package auth

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type contextKey string

const (
	UserSeqIDContext contextKey = "user-seq-context"

	UserSeqIDHeader = "X-User-Seq-ID"
)

type Authorizer interface {
	Authorize(ctx context.Context, route string) error
}

type AuthorizerInterceptor struct {
	tokens []string
}

// validateTokenFromHeader validates tokens in env config
func (ai *AuthorizerInterceptor) validateTokenFromHeader(md metadata.MD) bool {
	headers := md.Get("Authorization")

	if len(headers) == 0 {
		return false
	}

	key := strings.TrimPrefix(headers[0], "Bearer ")

	for _, k := range ai.tokens {
		if k == key {
			return true
		}
	}

	return false
}

func (ai *AuthorizerInterceptor) addUserTokenContext(ctx context.Context, seqID int) context.Context {
	return context.WithValue(ctx, UserSeqIDContext, seqID)
}

func (ai *AuthorizerInterceptor) getUserSeqID(md metadata.MD) (seq int, ok bool) {
	arr := md.Get(UserSeqIDHeader)

	if len(arr) == 0 {
		return 0, false
	}

	seq, err := strconv.Atoi(arr[0])

	if err != nil {
		return 0, false
	}

	return
}

// UnaryServerInterceptor handles authentication for each call
func UnaryServerInterceptor(tokens []string) grpc.UnaryServerInterceptor {
	ai := &AuthorizerInterceptor{
		tokens: tokens,
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)

		if !ok {
			return nil, status.Error(codes.PermissionDenied, "unauthorized: metadata missing")
		}

		if !ai.validateTokenFromHeader(md) {
			return nil, status.Error(codes.PermissionDenied, "unauthorized")
		}

		seq, ok := ai.getUserSeqID(md)

		if ok {
			ctx = ai.addUserTokenContext(ctx, seq)
		}

		srv, ok := info.Server.(Authorizer)
		if !ok {
			return nil, fmt.Errorf("each service should implement an authorization")
		}

		if err := srv.Authorize(ctx, info.FullMethod); err != nil {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}

		return handler(ctx, req)
	}
}

// GetValuesFromContext returns user and token stored in context
func GetValuesFromContext(ctx context.Context) (seqID int, ok bool) {
	v := ctx.Value(UserSeqIDContext)

	if v == nil {
		return 0, false
	}

	seqID, ok = v.(int)

	return
}
