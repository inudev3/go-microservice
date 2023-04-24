package main

import (
	protos "currency/protos/currency"
	"currency/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func main(){
	log := hclog.Default()
	gs :=grpc.NewServer()
	cs := server.Newcurrency(log)
	protos.RegisterCurrencyServer(gs, cs)
	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")
	if err!=nil{
		log.Error("unable to listen", "error", err)
		os.Exit(1)
	}
	gs.Serve(l)
}
