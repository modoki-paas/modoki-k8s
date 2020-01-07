package handler

import (
	"context"

	modoki "github.com/modoki-paas/modoki-k8s/api"
)

type UserOrgServer struct {
	Context *ServerContext
}

func (s *UserOrgServer) UserAdd(ctx context.Context, in *modoki.UserAddRequest) (*modoki.UserAddResponse, error) {
	return s.Context.UserOrgClient.UserAdd(ctx, in)
}

func (s *UserOrgServer) UserDelete(ctx context.Context, in *modoki.UserDeleteRequest) (*modoki.UserDeleteResponse, error) {
	return s.Context.UserOrgClient.UserDelete(ctx, in)
}

func (s *UserOrgServer) UserFindByID(ctx context.Context, in *modoki.UserFindByIDRequest) (*modoki.UserFindByIDResponse, error) {
	return s.Context.UserOrgClient.UserFindByID(ctx, in)
}

func (s *UserOrgServer) OrganizationAdd(ctx context.Context, in *modoki.OrganizationAddRequest) (*modoki.OrganizationAddResponse, error) {
	return s.Context.UserOrgClient.OrganizationAdd(ctx, in)
}

func (s *UserOrgServer) OrganizationDelete(ctx context.Context, in *modoki.OrganizationDeleteRequest) (*modoki.OrganizationDeleteResponse, error) {
	return s.Context.UserOrgClient.OrganizationDelete(ctx, in)
}

func (s *UserOrgServer) OrganizationInviteUser(ctx context.Context, in *modoki.OrganizationInviteUserRequest) (*modoki.OrganizationInviteUserResponse, error) {
	return s.Context.UserOrgClient.OrganizationInviteUser(ctx, in)
}

func (s *UserOrgServer) OrganizationListUser(ctx context.Context, in *modoki.OrganizationListUserRequest) (*modoki.OrganizationListUserResponse, error) {
	return s.Context.UserOrgClient.OrganizationListUser(ctx, in)
}

func (s *UserOrgServer) GetRoleBinding(ctx context.Context, in *modoki.GetRoleBindingRequest) (*modoki.GetRoleBindingResponse, error) {
	return s.Context.UserOrgClient.GetRoleBinding(ctx, in)
}
