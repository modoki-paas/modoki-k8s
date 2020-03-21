package main

import (
	"log"
	"net"

	extauth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	_ "github.com/go-sql-driver/mysql"
	api "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/authserver/config"
	"github.com/modoki-paas/modoki-k8s/authserver/handler"
	"github.com/modoki-paas/modoki-k8s/pkg/auth"
	"google.golang.org/grpc"
)

func initGRPCServer(sctx *handler.ServerContext) (*grpc.Server, error) {
	server := grpc.NewServer()

	api.RegisterAuthServer(server, &handler.AuthServer{Context: sctx})
	extauth.RegisterAuthorizationServer(server, &handler.ExtAuthZ{
		GA:      auth.NewGatewayAuthorizer(sctx.TokenClient, sctx.UserOrgClient),
		Context: sctx})

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
