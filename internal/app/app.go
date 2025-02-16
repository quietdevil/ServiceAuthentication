package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"serviceauth/internal/config"
	"serviceauth/internal/interceptor"
	"serviceauth/pkg/auth_v1"

	closer "github.com/quietdevil/Platform_common/pkg/closer"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider *serviceProvider
	gRPCserver      *grpc.Server
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
	lis, err := net.Listen("tcp", a.serviceProvider.GetGrpcConfig().Address())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server started on %v\n", a.serviceProvider.GrpcConfig.Address())
	if err = a.gRPCserver.Serve(lis); err != nil {
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
	a.gRPCserver = grpc.NewServer(grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(interceptor.ValidateInterceptor))
	reflection.Register(a.gRPCserver)
	auth_v1.RegisterAuthenticationServer(a.gRPCserver, a.serviceProvider.GetImplemintation(ctx))
	return nil
}
