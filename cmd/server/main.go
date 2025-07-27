package main

import (
	"log"
	"net"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/lypolix/grpc-calculator/pkg/adder"
	"github.com/lypolix/grpc-calculator/pkg/api"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"net/http"
)

func startGateway() {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	
	err := api.RegisterAdderHandlerFromEndpoint(
	  ctx, mux, "localhost:9090", opts,
	)
	if err != nil {
	  log.Fatal(err)
	}
  
	log.Println("HTTP Gateway started on :9090")
	http.ListenAndServe(":9090", mux)
  }
  

func main() {	
	go startGateway()
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
