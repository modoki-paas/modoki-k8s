package main

import (
	"net"
	"runtime/debug"

	_ "github.com/go-sql-driver/mysql"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	api "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/apiserver/config"
	"github.com/modoki-paas/modoki-k8s/apiserver/handler"
	"github.com/modoki-paas/modoki-k8s/internal/log"
	"github.com/modoki-paas/modoki-k8s/pkg/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func main() {
	logger := log.New()

	cfg, err := config.ReadConfig()

	if err != nil {
		panic(err)
	}

	sctx, err := handler.NewServerContext(cfg)

	if err != nil {
		logger.Fatalf("failed to initialize server context: %+v", err)
	}

	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		logger.Fatalf("failed to listen on :443: %+v", err)
	}

	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(func(p interface{}) error {
			logger.Printf("server paniced: %+v(trace: %s)", p, string(debug.Stack()))

			return grpc.Errorf(codes.Internal, "internal error")
		}),
	}

	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(opts...),
			auth.UnaryServerInterceptor(cfg.APIKeys),
			logger.UnaryInterceptor(),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(opts...),
			auth.StreamServerInterceptor(cfg.APIKeys),
			logger.StreamInterceptor(),
		),
	)

	api.RegisterTokenServer(server, &handler.TokenServer{Context: sctx})
	api.RegisterAppServer(server, &handler.AppServer{Context: sctx})
	api.RegisterUserOrgServer(server, &handler.UserOrgServer{Context: sctx})

	if err := server.Serve(listener); err != nil {
		logger.Fatalf("failed to start server on :80: %+v", err)
	}
}
