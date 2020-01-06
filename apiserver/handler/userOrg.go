package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	modoki "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/apiserver/store/users"
	"github.com/modoki-paas/modoki-k8s/internal/dbutil"
	"github.com/modoki-paas/modoki-k8s/internal/grpcutil"
	"github.com/modoki-paas/modoki-k8s/pkg/auth"
	"github.com/modoki-paas/modoki-k8s/pkg/rbac/permissions"
	"github.com/modoki-paas/modoki-k8s/pkg/rbac/roles"
	"github.com/modoki-paas/modoki-k8s/pkg/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserOrgServer struct {
	Context *ServerContext
}

var _ modoki.UserOrgServer = &UserOrgServer{}

func (s *UserOrgServer) UserAdd(ctx context.Context, req *modoki.UserAddRequest) (res *modoki.UserAddResponse, err error) {
	if err := auth.IsAuthorized(ctx, permissions.UserCreate); err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	err = dbutil.Transaction(ctx, s.Context.DB, func(tx *sqlx.Tx) error {
		userStore := users.NewUserStore(s.Context.DB)
		roleBindingStore := users.NewRoleBindingsStore(s.Context.DB)

		user := req.User

		systemRole := roles.FindSystemRole(user.SystemRoleName)

		if systemRole == nil {
			return status.Error(codes.InvalidArgument, "unknown system role")
		}

		seq, err := userStore.AddUser(user.UserId, user.Name, types.UserNormal, systemRole.Name)
		if err != nil {
			if err == users.ErrUserIDDuplicates {
				return status.Error(codes.AlreadyExists, "user id already exists")
			}

			return status.Error(codes.Internal, fmt.Sprintf("internal error: %s", err.Error()))
		}

		if _, err := roleBindingStore.RegisterRoleBinding(seq, seq, roles.Self.Name); err != nil {
			return status.Error(codes.Internal, fmt.Sprintf("internal error: %s", err.Error()))
		}

		res = &modoki.UserAddResponse{
			User: user,
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserOrgServer) UserDelete(_ context.Context, _ *modoki.UserDeleteRequest) (*modoki.UserDeleteResponse, error) {
	panic("not implemented")
}

func (s *UserOrgServer) UserFindByID(ctx context.Context, req *modoki.UserFindByIDRequest) (*modoki.UserFindByIDResponse, error) {
	if err := auth.IsAuthorized(ctx, permissions.UserGetAll); err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	userStore := users.NewUserStore(s.Context.DB)

	u, err := userStore.FindUserByID(req.UserId)

	if err != nil {
		if err == users.ErrUnknownUser {
			return nil, status.Error(codes.NotFound, "unknown user id")
		}

		return nil, status.Error(codes.Internal, fmt.Sprintf("internal error: %+v", err))
	}

	return &modoki.UserFindByIDResponse{
		User: &modoki.User{
			UserId:         u.ID,
			Name:           u.Name,
			SystemRoleName: u.SystemRole,
			CreatedAt:      grpcutil.GRPCTimestamp(u.CreatedAt),
			UpdatedAt:      grpcutil.GRPCTimestamp(u.UpdatedAt),
		},
	}, nil
}

func (s *UserOrgServer) OrganizationAdd(_ context.Context, _ *modoki.OrganizationAddRequest) (*modoki.OrganizationAddResponse, error) {
	panic("not implemented")
}

func (s *UserOrgServer) OrganizationDelete(_ context.Context, _ *modoki.OrganizationDeleteRequest) (*modoki.OrganizationDeleteResponse, error) {
	panic("not implemented")
}

func (s *UserOrgServer) OrganizationInviteUser(_ context.Context, _ *modoki.OrganizationInviteUserRequest) (*modoki.OrganizationInviteUserResponse, error) {
	panic("not implemented")
}

func (s *UserOrgServer) OrganizationListUser(_ context.Context, _ *modoki.OrganizationListUserRequest) (*modoki.OrganizationListUserResponse, error) {
	panic("not implemented")
}

func (s *UserOrgServer) GetRoleBinding(ctx context.Context, in *modoki.GetRoleBindingRequest) (*modoki.GetRoleBindingResponse, error) {
	if err := auth.IsAuthorized(ctx, permissions.UserOrgGetRoleBinding); err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	userStore := users.NewUserStore(s.Context.DB)
	rbStore := users.NewRoleBindingsStore(s.Context.DB)

	user, err := userStore.FindUserByID(in.UserId)

	if err != nil {
		if err == users.ErrUnknownUser {
			return nil, status.Error(codes.NotFound, "unknown user")
		}

		log.Println(err)
		return nil, status.Error(codes.Internal, "failed to find user")
	}

	targetUser, err := userStore.FindUserByID(in.TargetId)

	if err != nil {
		if err == users.ErrUnknownUser {
			return nil, status.Error(codes.NotFound, "unknown target user")
		}

		log.Println(err)
		return nil, status.Error(codes.Internal, "failed to find user")
	}

	rb, err := rbStore.GetRoleBinding(user.SeqID, targetUser.SeqID)

	if err != nil {
		if err == users.ErrUnknownRoleBinding {
			return nil, status.Error(codes.NotFound, "unknown role binding")
		}

		log.Println(err)
		return nil, status.Error(codes.Internal, "failed to find role binding")
	}

	return &modoki.GetRoleBindingResponse{
		Role: rb.Name,
	}, nil
}
