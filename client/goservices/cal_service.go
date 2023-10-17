package goservices

import (
	"context"
	"fmt"
)

type CalService interface {
	Hello(name string) error
}

type calService struct {
	calculateClient CalculateClient
}

func NewCalService(calculateClient CalculateClient) CalService {
	return calService{calculateClient}
}

func (bese calService) Hello(name string) error {
	req := HelloRequest{
		Name: name,
	}
	res, err := bese.calculateClient.Hello(context.Background(), &req)
	if err != nil {
		return err
	}
	fmt.Printf("Service :สวัสดีจ้า : \n")
	fmt.Printf("Req : %v\n", req.Name)
	fmt.Printf("Res : %v\n", res.Result)

	return nil
}
