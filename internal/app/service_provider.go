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

type ServiceProvider struct {
	pgConfig              config.PgConfig
	grpcConfig            config.GRPCConfig
	authenticationConfig  *config.AuthenticationConfig
	openSSlConfig         config.OpensslConfig
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

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (s *ServiceProvider) OpensslConfig() config.OpensslConfig {
	if s.openSSlConfig == nil {
		conf, err := config.NewOpensslConfig()
		if err != nil {
			log.Fatal(err)
		}
		s.openSSlConfig = conf
	}
	return s.openSSlConfig
}

func (s *ServiceProvider) PgConfig() config.PgConfig {
	if s.pgConfig == nil {
		pgConf, err := config.NewPgConfig()
		if err != nil {
			log.Fatal(err)
		}

		s.pgConfig = pgConf
	}
	return s.pgConfig
}

func (s *ServiceProvider) GrpcConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		grpc, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatal(err)
		}
		s.grpcConfig = grpc

	}
	return s.grpcConfig
}

func (s *ServiceProvider) AuthConfig() *config.AuthenticationConfig {
	if s.authenticationConfig == nil {
		conf, err := config.NewAuthConfig()
		if err != nil {
			log.Fatal(err)
		}
		s.authenticationConfig = conf
	}
	return s.authenticationConfig
}

func (s *ServiceProvider) ClientDb(ctx context.Context) db.Client {
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

func (s *ServiceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.repositoryUser == nil {
		rep := repos_user.NewRepository(s.ClientDb(ctx))
		s.repositoryUser = rep
	}
	return s.repositoryUser

}

func (s *ServiceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		man := transaction.NewManager(s.ClientDb(ctx).DB())
		s.txManager = man
	}
	return s.txManager
}

func (s *ServiceProvider) Logs(ctx context.Context) repository.Logger {
	if s.logger == nil {
		l := logs.NewLogs(s.ClientDb(ctx))
		s.logger = l
	}
	return s.logger
}

func (s *ServiceProvider) UserService(ctx context.Context) service.UserService {
	if s.serviceUser == nil {
		serviceU := serv_user.NewService(s.UserRepository(ctx), s.TxManager(ctx), s.Logs(ctx))
		s.serviceUser = serviceU
	}
	return s.serviceUser
}

func (s *ServiceProvider) ImplementationUser(ctx context.Context) *api_user.Implementation {
	if s.implementation == nil {
		impl := api_user.NewImplementation(s.UserService(ctx))
		s.implementation = impl
	}
	return s.implementation
}

func (s *ServiceProvider) ImplementationAuth(cxt context.Context) *api_auth.ImplementationAuthentication {
	if s.implementationAuth == nil {
		impl := api_auth.NewImplementationAuthentication(s.ServiceAuth(cxt))
		s.implementationAuth = impl
	}
	return s.implementationAuth
}

func (s *ServiceProvider) ImplementationAccess(cxt context.Context) *api_access.ImplementationAccess {
	if s.implementationAccess == nil {
		impl := api_access.NewImplementationAccess(s.ServiceAccess(cxt))
		s.implementationAccess = impl
	}
	return s.implementationAccess
}
func (s *ServiceProvider) AccessRepository(ctx context.Context) repository.AccessRepository {
	if s.repositoryAccess == nil {
		reposAccess := repos_access.NewAccessRepository(s.ClientDb(ctx))
		s.repositoryAccess = reposAccess
	}
	return s.repositoryAccess
}

func (s *ServiceProvider) AuthenticateRepository(ctx context.Context) repository.AuthenticationRepository {
	if s.repositoryAuth == nil {
		reposAuth := repos_auth.NewAuthRepos(s.ClientDb(ctx))
		s.repositoryAuth = reposAuth
	}
	return s.repositoryAuth
}

func (s *ServiceProvider) ServiceAccess(ctx context.Context) service.AccessService {
	if s.serviceAccess == nil {
		serv := access.NewAccessService(s.AccessRepository(ctx), s.AuthConfig())
		s.serviceAccess = serv
	}
	return s.serviceAccess
}

func (s *ServiceProvider) ServiceAuth(ctx context.Context) service.AuthenticationService {
	if s.serviceAuthentication == nil {
		serv := authentication.NewAuthenticationService(s.AuthenticateRepository(ctx), s.AuthConfig())
		s.serviceAuthentication = serv
	}
	return s.serviceAuthentication
}
