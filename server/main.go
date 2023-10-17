package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"server/goservices"
)

func main() {
	s := grpc.NewServer()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	goservices.RegisterCalculateServer(s, goservices.NewCalServer())

	fmt.Println("gRPC RUN : 50051")

	err = s.Serve(listener)
	if err != nil {

		log.Fatal(err)
	}
}
