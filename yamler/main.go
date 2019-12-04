package main

import (
	"log"
	"net"

	api "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/yamler/handler"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("failed to listen on :80: %v", err)
	}

	server := grpc.NewServer()
	api.RegisterGeneratorServer(server, &handler.Handler{})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to start server on :80: %v", err)
	}
}
