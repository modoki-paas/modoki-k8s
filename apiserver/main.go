package main

import (
	"flag"
	"log"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	api "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/apiserver/config"
	"github.com/modoki-paas/modoki-k8s/apiserver/handler"
	"github.com/modoki-paas/modoki-k8s/pkg/auth"
	"google.golang.org/grpc"
)

type commandArg struct {
	Config string
	Help   bool
}

func (arg *commandArg) Init() {
	flag.BoolVar(&arg.Help, "help", false, "show usage")
	flag.StringVar(&arg.Config, "config", "/etc/modoki/apiserver.yaml", "path to config file")

	flag.Parse()

	if arg.Help {
		flag.Usage()

		os.Exit(1)
	}
}

func main() {
	carg := &commandArg{}
	carg.Init()

	cfg, err := config.ReadConfig(carg.Config)

	if err != nil {
		panic(err)
	}

	envCfg, err := config.ReadEnv()

	if err != nil {
		panic(err)
	}

	sctx, err := handler.NewServerContext(cfg, envCfg)

	if err != nil {
		log.Fatalf("failed to initialize server context: %+v", err)
	}

	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		log.Fatalf("failed to listen on :443: %+v", err)
	}

	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			auth.UnaryServerInterceptor(envCfg.APIKeys),
		),
		grpc_middleware.WithStreamServerChain(
			auth.StreamServerInterceptor(envCfg.APIKeys),
		),
	)

	api.RegisterAppServer(server, &handler.AppServer{Context: sctx})
	api.RegisterUserOrgServer(server, &handler.UserOrgServer{Context: sctx})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to start server on :80: %+v", err)
	}
}
