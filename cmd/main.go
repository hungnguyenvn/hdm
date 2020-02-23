package main

import (
	"hdm"
	"hdm/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	gRPCAddr = ":8081"
)

func main() {
	errChan := make(chan error)
	go runHTTPServer(errChan)
	go runGRPCServer(errChan)

	log.Fatal(<-errChan)
}

func runHTTPServer(ch chan error) {
	server := hdm.NewHTTPServer()
	log.Println("Start to HTTP server")
	ch <- server.ListenAndServe()
}

func runGRPCServer(ch chan error) {
	handler := hdm.NewGRPCHandler()

	lis, err := net.Listen("tcp", gRPCAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHDMServer(grpcServer, handler)

	conn := hdm.ConnectGRPC(ch)
	log.Println("Connect to GRPC server")
	defer conn.Close()
	ch <- grpcServer.Serve(lis)
}
