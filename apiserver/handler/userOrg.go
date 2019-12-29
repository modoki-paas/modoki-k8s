package handler

import (
	"context"
	"fmt"

	modoki "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/apiserver/store/users"
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

func (s *UserOrgServer) UserAdd(ctx context.Context, req *modoki.UserAddRequest) (*modoki.UserAddResponse, error) {
	if err := auth.IsAuthorized(ctx, permissions.UserCreate); err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	userStore := users.NewUserStore(s.Context.DB)

	user := req.User

	systemRole := roles.FindSystemRole(user.RoleName)

	if systemRole == nil {
		return nil, status.Error(codes.InvalidArgument, "unknown system role")
	}

	_, err := userStore.AddUser(user.UserId, user.Name, types.UserNormal, systemRole.Name)
	if err != nil {
		if err == users.ErrUserIDDuplicates {
			return nil, status.Error(codes.AlreadyExists, "user id already exists")
		}

		return nil, status.Error(codes.Internal, fmt.Sprintf("internal error: %s", err.Error()))
	}

	return &modoki.UserAddResponse{
		User: user,
	}, nil
}

func (s *UserOrgServer) UserDelete(_ context.Context, _ *modoki.UserDeleteRequest) (*modoki.UserDeleteResponse, error) {
	panic("not implemented")
}

func (s *UserOrgServer) UserFindByID(_ context.Context, _ *modoki.UserFindByIDRequest) (*modoki.UserFindByIDResponse, error) {
	panic("not implemented")
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
