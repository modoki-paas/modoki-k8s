package grpcutil

import (
	"github.com/modoki-paas/modoki-k8s/pkg/auth"
	"github.com/modoki-paas/modoki-k8s/pkg/rbac/roles"
	"google.golang.org/grpc"
)

type GRPCDialer struct {
	token string
}

func NewGRPCDialer(token string) *GRPCDialer {
	return &GRPCDialer{
		token: token,
	}
}

func (gd *GRPCDialer) Dial(endpoint string, insecure bool) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{}

	if insecure {
		opts = append(opts, grpc.WithInsecure())
	}

	opts = append(
		opts,
		grpc.WithChainStreamInterceptor(
			auth.StreamClientInterceptor(gd.token),
		),
		grpc.WithChainUnaryInterceptor(
			auth.UnaryClientInterceptor(gd.token),
		),
	)

	return grpc.Dial(endpoint, opts...)
}

func (gd *GRPCDialer) DialAs(endpoint string, insecure bool, userID string, systemRole *roles.SystemRole) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{}

	if insecure {
		opts = append(opts, grpc.WithInsecure())
	}

	opts = append(
		opts,
		grpc.WithChainStreamInterceptor(
			auth.PerformerOverwritingStreamClientInterceptor(userID, systemRole),
			auth.StreamClientInterceptor(gd.token),
		),
		grpc.WithChainUnaryInterceptor(
			auth.PerformerOverwritingUnaryClientInterceptor(userID, systemRole),
			auth.UnaryClientInterceptor(gd.token),
		),
	)

	return grpc.Dial(endpoint, opts...)
}
