package main

import (
	"flag"
	"log"
	"net"
	"os"

	extauth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	_ "github.com/go-sql-driver/mysql"
	api "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/authserver/config"
	"github.com/modoki-paas/modoki-k8s/authserver/handler"
	"google.golang.org/grpc"
)

type commandArg struct {
	Config string
	Help   bool
}

func (arg *commandArg) Init() {
	flag.BoolVar(&arg.Help, "help", false, "show usage")
	flag.StringVar(&arg.Config, "config", "/etc/modoki/authserver.yaml", "path to config file")

	flag.Parse()

	if arg.Help {
		flag.Usage()

		os.Exit(1)
	}
}

func initGRPCServer(sctx *handler.ServerContext) (*grpc.Server, error) {
	cfg := sctx.Config

	server := grpc.NewServer()

	api.RegisterUserOrgServer(server, &handler.UserOrgServer{Context: sctx})
	api.RegisterTokenServer(server, &handler.TokenServer{Context: sctx})
	api.RegisterAppServer(server, &handler.AppServer{Context: sctx})
	api.RegisterAuthServer(server, &handler.AuthServer{Context: sctx})
	extauth.RegisterAuthorizationServer(server, &handler.ExtAuthZ{Context: sctx})

	return server, nil
}

func main() {

	carg := &commandArg{}
	carg.Init()

	cfg, err := config.ReadConfig(carg.Config)

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
