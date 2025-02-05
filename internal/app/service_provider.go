package app

import (
	"context"
	"log"
	api "serviceauth/internal/api/user"
	"serviceauth/internal/closer"
	"serviceauth/internal/config"
	"serviceauth/internal/repository"
	repos "serviceauth/internal/repository/user"
	"serviceauth/internal/service"
	serv "serviceauth/internal/service/user"

	"github.com/jackc/pgx/v5/pgxpool"
)

type serviceProvider struct {
	PgConfig       *config.PGConfig
	GrpcConfig     config.GRPCConfig
	Pool           *pgxpool.Pool
	Repository     repository.UserRepository
	Service        service.UserService
	Implemintation *api.Implementation
}

func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GetPgConfig() *config.PGConfig {
	if s.PgConfig == nil {
		pg, err := config.NewPgConfig()
		if err != nil {
			log.Fatal(err)
		}

		s.PgConfig = pg
	}
	return s.PgConfig
}

func (s *serviceProvider) GetGrpcConfig() config.GRPCConfig {
	if s.GrpcConfig == nil {
		grpc, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatal(err)
		}
		s.GrpcConfig = grpc

	}
	return s.GrpcConfig
}

func (s *serviceProvider) GetPool(ctx context.Context) *pgxpool.Pool {
	if s.Pool == nil {
		pg, err := pgxpool.New(ctx, s.GetPgConfig().GetDSN())
		if err != nil {
			log.Fatal(err)
		}
		if err = pg.Ping(ctx); err != nil {
			log.Fatal(err)
		}
		closer.Add(func() error {
			pg.Close()
			return nil
		})
		s.Pool = pg

	}
	return s.Pool
}

func (s *serviceProvider) GetRepository(ctx context.Context) repository.UserRepository {
	if s.Repository == nil {
		rep := repos.NewRepository(s.GetPool(ctx))
		s.Repository = rep
	}
	return s.Repository

}

func (s *serviceProvider) GetService(ctx context.Context) service.UserService {
	if s.Service == nil {
		service := serv.NewService(s.GetRepository(ctx))
		s.Service = service
	}
	return s.Service
}

func (s *serviceProvider) GetImplemintation(ctx context.Context) *api.Implementation {
	if s.Implemintation == nil {
		impl := api.NewImplementation(s.GetService(ctx))
		s.Implemintation = impl
	}
	return s.Implemintation
}
