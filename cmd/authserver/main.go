package main

import (
	"context"
	"log"
	"net"
	desc "serviceauth/pkg/auth_v1"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	desc.UnimplementedAuthenticationServer
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	return &desc.GetResponse{
		Id:        req.GetId(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
		Role:      *desc.Enum_admin.Enum(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterAuthenticationServer(s, &server{})

	log.Printf("server listing %v", lis.Addr().Network())

	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
