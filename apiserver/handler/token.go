package handler

import (
	"context"

	"github.com/jmoiron/sqlx"
	modoki "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/apiserver/store/tokens"
	"github.com/modoki-paas/modoki-k8s/internal/dbutil"
	"github.com/modoki-paas/modoki-k8s/internal/log"
	"github.com/modoki-paas/modoki-k8s/pkg/auth"
	"github.com/modoki-paas/modoki-k8s/pkg/rbac/permissions"
	"github.com/modoki-paas/modoki-k8s/pkg/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TokenServer struct {
	Context *ServerContext
}

var _ modoki.TokenServer = &TokenServer{}

func (s *TokenServer) IssueToken(ctx context.Context, req *modoki.IssueTokenRequest) (resp *modoki.IssueTokenResponse, err error) {
	if err := auth.IsAuthorized(ctx, permissions.TokenIssue); err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}
	logger := log.Extract(ctx)

	err = dbutil.Transaction(ctx, s.Context.DB, func(tx *sqlx.Tx) error {
		store := tokens.NewTokenStore(tx)

		_, tk, err := store.AddToken(&types.Token{
			Owner:  auth.GetTargetIDContext(ctx),
			Author: auth.GetUserIDContext(ctx),
			ID:     req.Id,
		})

		if err != nil {
			if err == tokens.ErrTokenIDDuplicates {
				return status.Error(codes.AlreadyExists, "the id already exists")
			}

			logger.Errorf("failed to generate token: %+v", err)
			return status.Error(codes.Internal, "failed to issue token")
		}

		resp = &modoki.IssueTokenResponse{
			Token: tk,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *TokenServer) ValidateToken(ctx context.Context, req *modoki.ValidateTokenRequest) (resp *modoki.ValidateTokenResponse, err error) {
	if err := auth.IsAuthorized(ctx, permissions.TokenValidate); err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}
	logger := log.Extract(ctx)

	err = dbutil.Transaction(ctx, s.Context.DB, func(tx *sqlx.Tx) error {
		store := tokens.NewTokenStore(tx)

		tk, err := store.FindTokenByToken(req.Token)

		if err != nil {
			if err == tokens.ErrUnknownToken {
				return status.Error(codes.NotFound, "unknown token")
			}

			logger.Errorf("failed to validate token: %+v", err)
			return status.Error(codes.Internal, "failed to find token")
		}

		resp = &modoki.ValidateTokenResponse{
			Id:        tk.ID,
			UserId:    tk.Owner,
			CreatedBy: tk.Author,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}
