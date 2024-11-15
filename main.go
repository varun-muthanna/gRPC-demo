package main

import (
	"context"
	"log"
	"net"

	"github.com/varun-muthanna/grpc/invoicer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Invoice struct {
	invoicer.UnimplementedInvoicerServer
}

func (i *Invoice) DefaultInvoicer(context.Context, *invoicer.Content) (*invoicer.Result, error) {
	b := []byte("Hello world")
	arr := []int64{1, 2, 3, 4}

	return &invoicer.Result{
		B:   b,
		Arr: arr,
	}, nil

}

func main() {

	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("Error establishing connection %s", err)
	}

	s := grpc.NewServer()
	service := &Invoice{}

	reflection.Register(s)

	invoicer.RegisterInvoicerServer(s, service)
	err1 := s.Serve(lis)

	if err1 != nil {
		log.Fatalf("Error establishing gRPC connection %s", err)
	}
}
