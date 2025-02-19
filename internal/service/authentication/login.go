package authentication

import (
	"context"
	"errors"
	"github.com/quietdevil/ServiceAuthentication/internal/model"
	"github.com/quietdevil/ServiceAuthentication/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"time"
)

const (
	AccessSecretKey        = "accessSecret"
	RefreshSecretKey       = "qwerty"
	refreshTokenExpiration = 60 * time.Minute
	accessTokenExpiration  = 5 * time.Minute
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

	refreshToken, err := utils.GenerateToken(userModel.UserInfo, []byte(RefreshSecretKey), refreshTokenExpiration)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return refreshToken, nil
}
