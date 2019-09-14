package main

import (
	"flag"
	"log"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	api "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/daemon/handler"
	"github.com/modoki-paas/modoki-k8s/daemon/store"
	"google.golang.org/grpc"
)

type commandArg struct {
	DSN  string
	Help bool
}

func (arg *commandArg) Init() {
	flag.StringVar(&arg.DSN, "db", "", "database source name")
	flag.BoolVar(&arg.Help, "help", false, "show usage")

	flag.Parse()

	if arg.Help {
		flag.Usage()

		os.Exit(1)
	}
}

func main() {
	carg := &commandArg{}
	carg.Init()

	d, err := sqlx.Open("mysql", carg.DSN)

	if err != nil {
		panic(err)
	}

	db := store.NewDB(d)

	sctx := &handler.ServerContext{
		DB: db,
	}

	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("failed to listen on :80: %v", err)
	}
	server := grpc.NewServer()
	api.RegisterServiceServer(server, &handler.ServiceServer{Context: sctx})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to start server on :80: %v", err)
	}
}
