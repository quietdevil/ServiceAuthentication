package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/quietdevil/ServiceAuthentication/pkg/access_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
)

var accessToken *string = flag.String("a", "", "")

func main() {
	flag.Parse()
	ctx := context.Background()

	mapC := make(map[string]string)
	mapC["Authorization"] = "Bearer " + *accessToken
	md := metadata.New(mapC)

	ctx = metadata.NewOutgoingContext(ctx, md)

	conn, err := grpc.NewClient("localhost:50000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := access_v1.NewAccessV1Client(conn)
	_, err = client.Check(ctx, &access_v1.CheckRequest{EndpointAddress: "/auth_v1.AuthenticationUserV1/Get"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Success")
}
