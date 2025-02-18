package authentication

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"serviceauth/internal/model"
	"serviceauth/internal/utils"
	"time"
)

const (
	AccessSecretKey        = "accessSecret"
	RefreshSecretKey       = "qwerty"
	refreshTokenExpiration = 60 * time.Minute
	accessTokenExpiration  = 5 * time.Minute
)

func (a *AuthenticationService) Login(ctx context.Context, login *model.UserLogin) (string, error) {
	//Ходит в репо слой за юзером
	userModel, err := a.reposAccess.GetUserByUsername(ctx, login.Username)
	if err != nil {
		return "", err
	}
	//Проверяет пароли
	if login.Password != userModel.UserInfo.Password {
		return "", status.Error(codes.PermissionDenied, "password does not match")
	}
	userModel.UserInfo.Role = "admin"
	refreshToken, err := utils.GenerateToken(userModel.UserInfo, []byte(RefreshSecretKey), refreshTokenExpiration)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return refreshToken, nil
}
