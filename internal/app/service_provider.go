package app

import (
	"context"
	api_access "github.com/quietdevil/ServiceAuthentication/internal/api/access"
	api_auth "github.com/quietdevil/ServiceAuthentication/internal/api/authentication"
	api_user "github.com/quietdevil/ServiceAuthentication/internal/api/user"
	"github.com/quietdevil/ServiceAuthentication/internal/config"
	"github.com/quietdevil/ServiceAuthentication/internal/repository"
	repos_access "github.com/quietdevil/ServiceAuthentication/internal/repository/access"
	repos_auth "github.com/quietdevil/ServiceAuthentication/internal/repository/authentication"
	"github.com/quietdevil/ServiceAuthentication/internal/repository/logs"
	repos_user "github.com/quietdevil/ServiceAuthentication/internal/repository/user"
	"github.com/quietdevil/ServiceAuthentication/internal/service"
	"github.com/quietdevil/ServiceAuthentication/internal/service/access"
	"github.com/quietdevil/ServiceAuthentication/internal/service/authentication"
	serv_user "github.com/quietdevil/ServiceAuthentication/internal/service/user"
	"log"

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
