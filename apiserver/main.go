package main

import (
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	api "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/apiserver/config"
	"github.com/modoki-paas/modoki-k8s/apiserver/handler"
	"github.com/modoki-paas/modoki-k8s/pkg/auth"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.ReadConfig()

	if err != nil {
		panic(err)
	}

	sctx, err := handler.NewServerContext(cfg)

	if err != nil {
		log.Fatalf("failed to initialize server context: %+v", err)
	}

	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		log.Fatalf("failed to listen on :443: %+v", err)
	}

	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			auth.UnaryServerInterceptor(cfg.APIKeys),
		),
		grpc_middleware.WithStreamServerChain(
			auth.StreamServerInterceptor(cfg.APIKeys),
		),
	)

	api.RegisterTokenServer(server, &handler.TokenServer{Context: sctx})
	api.RegisterAppServer(server, &handler.AppServer{Context: sctx})
	api.RegisterUserOrgServer(server, &handler.UserOrgServer{Context: sctx})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to start server on :80: %+v", err)
	}
}
