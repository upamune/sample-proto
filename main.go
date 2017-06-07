package main

import (
	"google.golang.org/grpc"
	"net"
	"log"
	"github.com/upamune/sample-proto/proto"
)

func main() {

	port := "50051"
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	srv := grpc.NewServer()
	proto.RegisterUserServiceServer(srv, &UserServer{})

	log.Println("Starting server port:", port)
	if err := srv.Serve(l); err != nil {
		log.Fatal(err)
	}
}

