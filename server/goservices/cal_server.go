package goservices

import (
	context "context"
	"fmt"
)

type calServer struct {
}

func NewCalServer() CalculateServer {
	return calServer{}
}

func (calServer) mustEmbedUnimplementedCalculateServer() {}

func (calServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	result := fmt.Sprintf("สวัสดี[Service] %v", req.Name)
	res := HelloResponse{
		Result: result,
	}
	return &res, nil
}
