package main

import (
	"flag"
	"log"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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

func main() {
	sctx := &handler.ServerContext{}

	carg := &commandArg{}
	carg.Init()

	cfg, err := config.ReadConfig(carg.Config)

	if err != nil {
		panic(err)
	}

	sctx.Config = cfg

	d, err := sqlx.Open("mysql", cfg.DB)

	if err != nil {
		panic(err)
	}

	sctx.DB = d

	listener, err := net.Listen("tcp", ":443")
	if err != nil {
		log.Fatalf("failed to listen on :443: %v", err)
	}

	server := grpc.NewServer()

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to start server on :80: %v", err)
	}
}
