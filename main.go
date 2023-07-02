package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	protos "github.com/DShaKi/Currency-API/protos/currency"
	"github.com/DShaKi/Currency-API/server"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	cs := server.NewCurrency(log)

	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Error("Staring the gRPC server encountered an error: ", err)
		os.Exit(1)
	}

	gs.Serve(lis)
}
