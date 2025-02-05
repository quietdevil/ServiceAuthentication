package config

import (
	"errors"
	"net"
	"os"
)

type GRPCConfig interface {
	Address() string
}

type GrpcConfig struct {
	host string
	port string
}

func NewGRPCConfig() (GRPCConfig, error) {
	host := os.Getenv("GRPC_HOST")
	if len(host) == 0 {
		return nil, errors.New("grpc host don't parse")
	}
	port := os.Getenv("GRPC_PORT")
	if len(port) == 0 {
		return nil, errors.New("grpc port don't parse")
	}
	return &GrpcConfig{
		host: host,
		port: port,
	}, nil

}

func (c GrpcConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
