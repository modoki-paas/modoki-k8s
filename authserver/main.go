package main

import (
	"net"
	"runtime/debug"

	extauth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	_ "github.com/go-sql-driver/mysql"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	api "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/authserver/config"
	"github.com/modoki-paas/modoki-k8s/authserver/handler"
	"github.com/modoki-paas/modoki-k8s/internal/log"
	"github.com/modoki-paas/modoki-k8s/pkg/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var logger *log.Logger

func recoveryFunc(p interface{}) error {
	logger.Printf("server paniced: %+v(trace: %s)", p, string(debug.Stack()))

	return grpc.Errorf(codes.Internal, "internal error")
}

func initGRPCServer(sctx *handler.ServerContext) (*grpc.Server, error) {

	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(recoveryFunc),
	}

	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(opts...),
			logger.UnaryInterceptor(),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(opts...),
			logger.StreamInterceptor(),
		),
	)

	api.RegisterAuthServer(server, &handler.AuthServer{Context: sctx})
	extauth.RegisterAuthorizationServer(server, &handler.ExtAuthZ{
		GA:      auth.NewGatewayAuthorizer(sctx.TokenClient, sctx.UserOrgClient),
		Context: sctx},
	)

	return server, nil
}

func main() {
	logger := log.New()

	cfg, err := config.ReadConfig()

	if err != nil {
		panic(err)
	}

	sctx, err := handler.NewServerContext(cfg)

	if err != nil {
		logger.Fatalf("failed to init server context: %+v", err)
	}

	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		logger.Fatalf("failed to listen on %s: %+v", cfg.Address, err)
	}

	server, err := initGRPCServer(sctx)

	if err != nil {
		logger.Fatalf("failed to init grpc server: %+v", err)
	}

	if err := server.Serve(listener); err != nil {
		logger.Fatalf("failed to start server on :80: %v", err)
	}
}
