package app

import (
	"context"
	"fmt"
	"github.com/quietdevil/Platform_common/pkg/closer"
	"github.com/quietdevil/ServiceAuthentication/internal/config"
	"github.com/quietdevil/ServiceAuthentication/internal/interceptor"
	"github.com/quietdevil/ServiceAuthentication/pkg/access_v1"
	"github.com/quietdevil/ServiceAuthentication/pkg/auth_user_v1"
	"github.com/quietdevil/ServiceAuthentication/pkg/auth_v1"
	"google.golang.org/grpc/credentials"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider *ServiceProvider
	gRPCServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}
	if err := app.initDeps(ctx); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()
	return a.RunGRPCServer()
}

func (a *App) RunGRPCServer() error {
	lis, err := net.Listen("tcp", a.serviceProvider.GrpcConfig().Address())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server started on %v\n", a.serviceProvider.grpcConfig.Address())
	if err = a.gRPCServer.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	fn := []func(ctx context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGrpcServer,
	}

	for _, v := range fn {
		err := v(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	serProv := NewServiceProvider()
	a.serviceProvider = serProv
	return nil
}

func (a *App) initGrpcServer(ctx context.Context) error {
	transportCreds, err := credentials.NewServerTLSFromFile(a.serviceProvider.OpensslConfig().PathPem(), a.serviceProvider.OpensslConfig().PathKey())
	if err != nil {
		return err
	}
	a.gRPCServer = grpc.NewServer(grpc.Creds(transportCreds),
		grpc.UnaryInterceptor(interceptor.ValidateInterceptor))

	reflection.Register(a.gRPCServer)
	auth_user_v1.RegisterAuthenticationUserV1Server(a.gRPCServer, a.serviceProvider.ImplementationUser(ctx))
	auth_v1.RegisterAuthenticationV1Server(a.gRPCServer, a.serviceProvider.ImplementationAuth(ctx))
	access_v1.RegisterAccessV1Server(a.gRPCServer, a.serviceProvider.ImplementationAccess(ctx))
	return nil
}
