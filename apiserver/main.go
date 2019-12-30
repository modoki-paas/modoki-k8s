package main

import (
	"flag"
	"log"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
	api "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/apiserver/config"
	"github.com/modoki-paas/modoki-k8s/apiserver/handler"
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

	sctx, err := handler.NewServerContext(cfg)

	if err != nil {
		log.Fatalf("failed to initialize server context: %+v", err)
	}

	listener, err := net.Listen("tcp", ":443")
	if err != nil {
		log.Fatalf("failed to listen on :443: %+v", err)
	}

	server := grpc.NewServer()
	api.RegisterAppServer(server, &handler.AppServer{Context: sctx})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to start server on :80: %+v", err)
	}
}
