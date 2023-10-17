package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"client/goservices"
)

func main() {
	creds := insecure.NewCredentials()
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()

	CalculateClient := goservices.NewCalculateClient(cc)
	CalService := goservices.NewCalService(CalculateClient)

	err = CalService.Hello("ICESAMA 001")
	if err != nil {
		log.Fatal(err)
	}

}
