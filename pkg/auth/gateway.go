package auth

import (
	"context"

	api "github.com/modoki-paas/modoki-k8s/api"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// ErrUnauthenticated returns an error when authentication failed
	ErrUnauthenticated = xerrors.Errorf("unauthenticated")
)

type GatewayAuthorizer struct {
	tokenClient   api.TokenClient
	userOrgClient api.UserOrgClient
}

func NewGatewayAuthorizer(tokenClient api.TokenClient, userOrgClient api.UserOrgClient) *GatewayAuthorizer {
	return &GatewayAuthorizer{
		tokenClient:   tokenClient,
		userOrgClient: userOrgClient,
	}
}

// AuthenticatedMetadata represents data retrieved from the access token
type AuthenticatedMetadata struct {
	UserID string
	Roles  RoleBindings

	TargetID             string
	PermissionsForTarget map[string]struct{}
}

// GetAuthenticatedMetadata returns all metadata to authorize users
func (ai *GatewayAuthorizer) GetAuthenticatedMetadata(ctx context.Context, tk, targetID string) (*AuthenticatedMetadata, error) {
	vt, err := ai.tokenClient.ValidateToken(ctx, &api.ValidateTokenRequest{
		Token: tk,
	})

	if err != nil {
		if stat, ok := status.FromError(err); ok {
			switch stat.Code() {
			case codes.NotFound:
				return nil, ErrUnauthenticated
			}
		}

		return nil, xerrors.Errorf("failed to validate token: %w", err)
	}

	performer, err := ai.userOrgClient.UserFindByID(ctx, &api.UserFindByIDRequest{
		UserId: vt.UserId,
	})

	if err != nil {
		return nil, xerrors.Errorf("failed to validate token: %w", err)
	}

	if targetID == "" {
		targetID = performer.User.UserId
	}

	rb, err := ai.userOrgClient.GetRoleBinding(ctx, &api.GetRoleBindingRequest{
		UserId:   performer.GetUser().UserId,
		TargetId: targetID,
	})

	if err != nil {
		if stat, ok := status.FromError(err); !ok {
			return nil, xerrors.Errorf("failed to get role binding: %w", err)
		} else if stat.Code() != codes.NotFound {
			return nil, stat.Err()
		}
		// continue even if role binding is not found
	}

	roles := RoleBindings(map[string]string{
		"*": performer.User.SystemRoleName,
	})

	if rb != nil && rb.Role != "" {
		roles[targetID] = rb.Role
	}

	perms := getPermissions(roles, targetID)

	return &AuthenticatedMetadata{
		UserID: performer.User.UserId,
		Roles:  roles,

		TargetID:             targetID,
		PermissionsForTarget: perms,
	}, nil
}
