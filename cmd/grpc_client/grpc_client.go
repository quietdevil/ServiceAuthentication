package grpc_client

import (
	"github.com/quietdevil/ServiceAuthentication/pkg/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

func GrpcAuthClient(host, port, secretPath string) *auth_v1.AuthenticationV1Client {
	tramCreds, err := credentials.NewClientTLSFromFile(secretPath, "")
	conn, err := grpc.NewClient(net.JoinHostPort(host, port), grpc.WithTransportCredentials(tramCreds))
	if err != nil {
		log.Fatal(err)
	}
	clientAuth := auth_v1.NewAuthenticationV1Client(conn)
	return &clientAuth
}
