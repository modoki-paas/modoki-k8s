package main

import (
	"net"
	"runtime/debug"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	api "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/internal/k8s"
	"github.com/modoki-paas/modoki-k8s/internal/log"
	"github.com/modoki-paas/modoki-k8s/pkg/auth"
	"github.com/modoki-paas/modoki-k8s/yamler/config"
	"github.com/modoki-paas/modoki-k8s/yamler/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func main() {
	logger := log.New()

	cfg, err := config.ReadConfig()

	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		logger.Fatalf("failed to listen on %s: %v", cfg.Address, err)
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

	// in-cluster mode
	k8sClient, err := k8s.NewClient("")

	if err != nil {
		logger.Fatalf("failed to initialize kubernetes client: %+v", err)
	}

	api.RegisterGeneratorServer(
		server,
		&handler.Handler{
			Config:    cfg,
			K8sClient: k8sClient,
		},
	)

	if err := server.Serve(listener); err != nil {
		logger.Fatalf("failed to start server on :80: %v", err)
	}
}
