package auth

import (
	"context"
	"encoding/json"

	"github.com/modoki-paas/modoki-k8s/pkg/rbac/roles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ClientInterceptor struct {
	token string
}

func (ci *ClientInterceptor) addToken(md metadata.MD) {
	md.Set("Authorization", ci.token)
}

func (ci *ClientInterceptor) setUserID(ctx context.Context, md metadata.MD) {
	md.Set(UserIDHeader, GetUserIDContext(ctx))
}

func (ci *ClientInterceptor) setTargetID(ctx context.Context, md metadata.MD) {
	md.Set(UserIDHeader, GetTargetIDContext(ctx))
}

func (ci *ClientInterceptor) setRoles(ctx context.Context, md metadata.MD) {
	b, _ := json.Marshal(GetRolesContext(ctx))

	md.Set(RolesHeader, string(b))
}

func (ci *ClientInterceptor) wrapContext(ctx context.Context) context.Context {
	md, ok := metadata.FromOutgoingContext(ctx)

	if !ok {
		md = metadata.New(nil)
	}

	ci.addToken(md)
	ci.setUserID(ctx, md)
	ci.setTargetID(ctx, md)
	ci.setRoles(ctx, md)

	return metadata.NewOutgoingContext(ctx, md)
}

func UnaryClientInterceptor(token string) grpc.UnaryClientInterceptor {
	ci := &ClientInterceptor{
		token: token,
	}

	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx = ci.wrapContext(ctx)

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func StreamClientInterceptor(token string) grpc.StreamClientInterceptor {
	ci := &ClientInterceptor{
		token: token,
	}

	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		ctx = ci.wrapContext(ctx)

		return streamer(ctx, desc, cc, method, opts...)
	}
}

// PerformerOverwritingUnaryClientInterceptor calls other service explicitly as the specified user with system role
func PerformerOverwritingUnaryClientInterceptor(userID string, systemRole *roles.SystemRole) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		rb := RoleBindings(map[string]string{
			"*": systemRole.Name,
		})

		if ctx.Value(TargetIDContext) == nil {
			ctx = AddTargetIDContext(ctx, userID)
		}

		ctx = AddUserIDContext(ctx, userID)
		ctx = AddRolesContext(ctx, rb)
		ctx = AddPermissionsContext(ctx, getPermissions(rb, ""))

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// PerformerOverwritingStreamClientInterceptor calls other service explicitly as the specified user with system role
func PerformerOverwritingStreamClientInterceptor(userID string, systemRole *roles.SystemRole) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		rb := RoleBindings(map[string]string{
			"*": systemRole.Name,
		})

		if ctx.Value(TargetIDContext) == nil {
			ctx = AddTargetIDContext(ctx, userID)
		}

		ctx = AddUserIDContext(ctx, userID)
		ctx = AddRolesContext(ctx, rb)
		ctx = AddPermissionsContext(ctx, getPermissions(rb, ""))

		return streamer(ctx, desc, cc, method, opts...)
	}
}
