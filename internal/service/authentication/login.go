package authentication

import (
	"context"
	"errors"
	"github.com/quietdevil/ServiceAuthentication/internal/model"
	"github.com/quietdevil/ServiceAuthentication/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *AuthenticationService) Login(ctx context.Context, login *model.UserLogin) (string, error) {
	userModel, err := a.reposAccess.GetUserByUsername(ctx, login.Username)
	if err != nil {
		return "", err
	}

	err = utils.ComparePasswords(userModel.Password, login.Password)
	if err != nil {
		return "", status.Error(codes.PermissionDenied, "password does not match")
	}

	refreshToken, err := utils.GenerateToken(userModel.UserInfo, []byte(a.configAuth.RefreshSecretKey()), a.configAuth.RefreshTime())
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return refreshToken, nil
}
