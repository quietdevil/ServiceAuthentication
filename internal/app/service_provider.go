package app

import (
	"context"
	"log"
	api "serviceauth/internal/api/user"
	"serviceauth/internal/config"
	"serviceauth/internal/repository"
	"serviceauth/internal/repository/logs"
	repos "serviceauth/internal/repository/user"
	"serviceauth/internal/service"
	serv "serviceauth/internal/service/user"

	db "github.com/quietdevil/Platform_common/pkg/db"
	pg "github.com/quietdevil/Platform_common/pkg/db/pg"
	transaction "github.com/quietdevil/Platform_common/pkg/db/transaction"

	closer "github.com/quietdevil/Platform_common/pkg/closer"
)

type serviceProvider struct {
	PgConfig       *config.PGConfig
	GrpcConfig     config.GRPCConfig
	ClientDB       db.Client
	TxManager      db.TxManager
	Logger         repository.Logger
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

func (s *serviceProvider) GetClient(ctx context.Context) db.Client {
	if s.ClientDB == nil {
		clientdb, err := pg.NewDBClient(ctx, s.GetPgConfig().GetDSN())
		if err != nil {
			log.Fatal(err)
		}
		if err = clientdb.DB().Ping(ctx); err != nil {
			log.Fatal(err)
		}

		closer.Add(clientdb.Close)
		s.ClientDB = clientdb

	}
	return s.ClientDB
}

func (s *serviceProvider) GetRepository(ctx context.Context) repository.UserRepository {
	if s.Repository == nil {
		rep := repos.NewRepository(s.GetClient(ctx))
		s.Repository = rep
	}
	return s.Repository

}

func (s *serviceProvider) GetTxManager(ctx context.Context) db.TxManager {
	if s.TxManager == nil {
		man := transaction.NewManager(s.ClientDB.DB())
		s.TxManager = man
	}
	return s.TxManager
}

func (s *serviceProvider) Logs(ctx context.Context) repository.Logger {
	if s.Logger == nil {
		log := logs.NewLogs(s.GetClient(ctx))
		s.Logger = log
	}
	return s.Logger
}

func (s *serviceProvider) GetService(ctx context.Context) service.UserService {
	if s.Service == nil {
		service := serv.NewService(s.GetRepository(ctx), s.GetTxManager(ctx), s.Logs(ctx))
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
