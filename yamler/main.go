package main

import (
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	api "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/internal/k8s"
	"github.com/modoki-paas/modoki-k8s/pkg/auth"
	"github.com/modoki-paas/modoki-k8s/yamler/config"
	"github.com/modoki-paas/modoki-k8s/yamler/handler"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.ReadConfig()

	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", cfg.Address, err)
	}

	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			auth.UnaryServerInterceptor(cfg.APIKeys),
		),
		grpc_middleware.WithStreamServerChain(
			auth.StreamServerInterceptor(cfg.APIKeys),
		),
	)

	// in-cluster mode
	k8sClient, err := k8s.NewClient("")

	if err != nil {
		log.Fatalf("failed to initialize kubernetes client: %+v", err)
	}

	api.RegisterGeneratorServer(
		server,
		&handler.Handler{
			Config:    cfg,
			K8sClient: k8sClient,
		},
	)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to start server on :80: %v", err)
	}
}
