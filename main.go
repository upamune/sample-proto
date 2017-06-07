package main

import (
	"google.golang.org/grpc"
	"net"
	"log"
	"github.com/upamune/sample-proto/proto"
)

func main() {

	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	srv := grpc.NewServer()
	proto.RegisterUserServiceServer(srv, &UserServer{})

	if err := srv.Serve(l); err != nil {
		log.Fatal(err)
	}
}

