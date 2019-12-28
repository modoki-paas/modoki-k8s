package handler

import (
	"context"

	modoki "github.com/modoki-paas/modoki-k8s/api"
)

type UserOrgServer struct {
}

func (s *UserOrgServer) UserAdd(_ context.Context, _ *modoki.UserAddRequest) (*modoki.UserAddResponse, error) {
	panic("not implemented")
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
