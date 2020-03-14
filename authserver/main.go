package main

import (
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	api "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/authserver/config"
	"github.com/modoki-paas/modoki-k8s/authserver/handler"
	"github.com/modoki-paas/modoki-k8s/internal/connector"
	"github.com/modoki-paas/modoki-k8s/internal/grpcutil"
	"github.com/modoki-paas/modoki-k8s/pkg/auth"
	"github.com/modoki-paas/modoki-k8s/pkg/rbac/roles"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

func initGRPCServer(sctx *handler.ServerContext) (*grpc.Server, error) {
	cfg := sctx.Config

	dialer := grpcutil.NewGRPCDialer(cfg.APIKeys[0])

	dialer.UnaryClientInterceptors = append(
		[]grpc.UnaryClientInterceptor{auth.PerformerOverwritingUnaryClientInterceptor("authserver", roles.SystemAuth)},
		dialer.UnaryClientInterceptors...,
	)

	dialer.StreamClientInterceptors = append(
		[]grpc.StreamClientInterceptor{auth.PerformerOverwritingStreamClientInterceptor("authserver", roles.SystemAuth)},
		dialer.StreamClientInterceptors...,
	)

	connector := connector.NewConnector(dialer)

	userOrg, err := connector.ConnectUserOrgClient(cfg.Endpoints.UserOrg.Endpoint, cfg.Endpoints.UserOrg.Insecure)

	if err != nil {
		return nil, xerrors.Errorf("failed to initialize user/org client for gateway: %w", err)
	}

	token, err := connector.ConnectTokenClient(cfg.Endpoints.Token.Endpoint, cfg.Endpoints.Token.Insecure)

	if err != nil {
		return nil, xerrors.Errorf("failed to initialize user/org client for gateway: %w", err)
	}

	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			auth.UnaryGatewayServerInterceptor(token, userOrg),
		),
		grpc_middleware.WithStreamServerChain(
			auth.StreamGatewayServerInterceptor(token, userOrg),
		),
	)

	api.RegisterUserOrgServer(server, &handler.UserOrgServer{Context: sctx})
	api.RegisterTokenServer(server, &handler.TokenServer{Context: sctx})
	api.RegisterAppServer(server, &handler.AppServer{Context: sctx})
	api.RegisterAuthServer(server, &handler.AuthServer{Context: sctx})

	return server, nil
}

func main() {
	cfg, err := config.ReadConfig()

	if err != nil {
		panic(err)
	}

	sctx, err := handler.NewServerContext(cfg)

	if err != nil {
		log.Fatalf("failed to init server context: %+v", err)
	}

	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		log.Fatalf("failed to listen on %s: %+v", cfg.Address, err)
	}

	server, err := initGRPCServer(sctx)

	if err != nil {
		log.Fatalf("failed to init grpc server: %+v", err)
	}

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to start server on :80: %v", err)
	}
}
