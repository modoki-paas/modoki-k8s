package handler

import (
	"context"
	"fmt"

	"strings"

	"database/sql"

	"github.com/modoki-paas/modoki-k8s/apiserver/store"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type contextKey string

const (
	TokenContext contextKey = "token-context"
	UserContext  contextKey = "user-context"
)

type Authorizer interface {
	Authorize(ctx context.Context, route string) error
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

func (ai *AuthorizerInterceptor) addUserTokenContext(ctx context.Context) (context.Context, error) {
	var user *store.User
	var tk *store.Token

	token, err := ai.getTokenFromHeader(ctx)

	if err == nil {
		user, tk, err = ai.db.User().GetUserFromToken(token)

		if xerrors.Is(err, sql.ErrNoRows) {
			// unauthorized
			user = nil
			tk = nil
		} else {
			return nil, xerrors.Errorf("failed to get user and token from db: %v", err)
		}
	}

	ctx = context.WithValue(ctx, TokenContext, tk)
	ctx = context.WithValue(ctx, UserContext, user)

	return ctx, nil
}

// UnaryServerInterceptor はリクエストごとの認可を行う、unary サーバーインターセプターを返す。
func UnaryServerInterceptor(db *store.DB) grpc.UnaryServerInterceptor {
	ai := &AuthorizerInterceptor{db}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx, err := ai.addUserTokenContext(ctx)

		if err != nil {
			return nil, xerrors.Errorf("failed to add user and token data to context: %v", err)
		}

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

// GetValuesFromContext returns user and token stored in context
func GetValuesFromContext(ctx context.Context) (user *store.User, token *store.Token) {
	u := ctx.Value(UserContext)

	if u == nil {
		user = nil
	} else {
		user = u.(*store.User)
	}

	tk := ctx.Value(TokenContext)

	if tk != nil {
		token = tk.(*store.Token)
	}

	return
}
