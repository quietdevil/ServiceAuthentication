package app

import (
	"context"
	"log"
	api_access "serviceauth/internal/api/access"
	api_auth "serviceauth/internal/api/authentication"
	api_user "serviceauth/internal/api/user"
	"serviceauth/internal/config"
	"serviceauth/internal/repository"
	repos_access "serviceauth/internal/repository/access"
	repos_auth "serviceauth/internal/repository/authentication"
	"serviceauth/internal/repository/logs"
	repos_user "serviceauth/internal/repository/user"
	"serviceauth/internal/service"
	"serviceauth/internal/service/access"
	"serviceauth/internal/service/authentication"
	serv_user "serviceauth/internal/service/user"

	db "github.com/quietdevil/Platform_common/pkg/db"
	pg "github.com/quietdevil/Platform_common/pkg/db/pg"
	transaction "github.com/quietdevil/Platform_common/pkg/db/transaction"

	closer "github.com/quietdevil/Platform_common/pkg/closer"
)

type serviceProvider struct {
	pgConfig              config.PgConfig
	grpcConfig            config.GRPCConfig
	clientDB              db.Client
	txManager             db.TxManager
	logger                repository.Logger
	repositoryUser        repository.UserRepository
	repositoryAccess      repository.AccessRepository
	repositoryAuth        repository.AuthenticationRepository
	serviceUser           service.UserService
	serviceAccess         service.AccessService
	serviceAuthentication service.AuthenticationService
	implementation        *api_user.Implementation
	implementationAuth    *api_auth.ImplementationAuthentication
	implementationAccess  *api_access.ImplementationAccess
}

func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PgConfig() config.PgConfig {
	if s.pgConfig == nil {
		pgConf, err := config.NewPgConfig()
		if err != nil {
			log.Fatal(err)
		}

		s.pgConfig = pgConf
	}
	return s.pgConfig
}

func (s *serviceProvider) GrpcConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		grpc, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatal(err)
		}
		s.grpcConfig = grpc

	}
	return s.grpcConfig
}

func (s *serviceProvider) ClientDb(ctx context.Context) db.Client {
	if s.clientDB == nil {
		clientDb, err := pg.NewDBClient(ctx, s.PgConfig().DSN())
		if err != nil {
			log.Fatal(err)
		}
		if err = clientDb.DB().Ping(ctx); err != nil {
			log.Fatal(err)
		}

		closer.Add(clientDb.Close)
		s.clientDB = clientDb

	}
	return s.clientDB
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.repositoryUser == nil {
		rep := repos_user.NewRepository(s.ClientDb(ctx))
		s.repositoryUser = rep
	}
	return s.repositoryUser

}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		man := transaction.NewManager(s.ClientDb(ctx).DB())
		s.txManager = man
	}
	return s.txManager
}

func (s *serviceProvider) Logs(ctx context.Context) repository.Logger {
	if s.logger == nil {
		log := logs.NewLogs(s.ClientDb(ctx))
		s.logger = log
	}
	return s.logger
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.serviceUser == nil {
		service := serv_user.NewService(s.UserRepository(ctx), s.TxManager(ctx), s.Logs(ctx))
		s.serviceUser = service
	}
	return s.serviceUser
}

func (s *serviceProvider) ImplementationUser(ctx context.Context) *api_user.Implementation {
	if s.implementation == nil {
		impl := api_user.NewImplementation(s.UserService(ctx))
		s.implementation = impl
	}
	return s.implementation
}

func (s *serviceProvider) ImplementationAuth(cxt context.Context) *api_auth.ImplementationAuthentication {
	if s.implementationAuth == nil {
		impl := api_auth.NewImplementationAuthentication(s.ServiceAuth(cxt))
		s.implementationAuth = impl
	}
	return s.implementationAuth
}

func (s *serviceProvider) ImplementationAccess(cxt context.Context) *api_access.ImplementationAccess {
	if s.implementationAccess == nil {
		impl := api_access.NewImplementationAccess(s.ServiceAccess(cxt))
		s.implementationAccess = impl
	}
	return s.implementationAccess
}
func (s *serviceProvider) AccessRepository(ctx context.Context) repository.AccessRepository {
	if s.repositoryAccess == nil {
		reposAccess := repos_access.NewAccessRepository(s.ClientDb(ctx))
		s.repositoryAccess = reposAccess
	}
	return s.repositoryAccess
}

func (s *serviceProvider) AuthenticateRepository(ctx context.Context) repository.AuthenticationRepository {
	if s.repositoryAuth == nil {
		reposAuth := repos_auth.NewAuthRepos(s.ClientDb(ctx))
		s.repositoryAuth = reposAuth
	}
	return s.repositoryAuth
}

func (s *serviceProvider) ServiceAccess(ctx context.Context) service.AccessService {
	if s.serviceAccess == nil {
		serv := access.NewAccessService(s.AccessRepository(ctx))
		s.serviceAccess = serv
	}
	return s.serviceAccess
}

func (s *serviceProvider) ServiceAuth(ctx context.Context) service.AuthenticationService {
	if s.serviceAuthentication == nil {
		serv := authentication.NewAuthenticationService(s.AuthenticateRepository(ctx))
		s.serviceAuthentication = serv
	}
	return s.serviceAuthentication
}
