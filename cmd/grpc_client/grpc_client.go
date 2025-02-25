package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/quietdevil/ServiceAuthentication/pkg/access_v1"
	"github.com/quietdevil/ServiceAuthentication/pkg/auth_user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

var accessToken *string = flag.String("a", "", "")

func main() {
	flag.Parse()
	ctx := context.Background()

	//mapC := make(map[string]string)
	//mapC["Authorization"] = "Bearer " + *accessToken
	//md := metadata.New(mapC)
	//
	//ctx = metadata.NewOutgoingContext(ctx, md)

	tramCreds, err := credentials.NewClientTLSFromFile("keys/service.pem", "")

	conn, err := grpc.NewClient("localhost:50000", grpc.WithTransportCredentials(tramCreds))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conn)

	clientAuth := access_v1.NewAccessV1Client(conn)
	clientUser := auth_user_v1.NewAuthenticationUserV1Client(conn)
	_, err = clientAuth.Check(ctx, &access_v1.CheckRequest{EndpointAddress: "/auth_v1.AuthenticationUserV1/Get"})
	if err != nil {
		log.Fatal(err)
	}

	_, err = clientUser.Get(ctx, &auth_user_v1.GetRequest{Id: int64(1)})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Success")
}
