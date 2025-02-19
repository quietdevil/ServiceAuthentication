package service

import (
	"context"
	"github.com/quietdevil/ServiceAuthentication/internal/model"
)

type UserService interface {
	Get(context.Context, int) (*model.User, error)
	Create(context.Context, *model.UserInfo) (int, error)
	Delete(context.Context, int) error
	Update(context.Context, *model.UserUpdate) error
}

type AccessService interface {
	Check(ctx context.Context, endpoint string) error
}

type AuthenticationService interface {
	Login(ctx context.Context, login *model.UserLogin) (string, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
	GetRefreshToken(ctx context.Context, refreshToken string) (string, error)
}
