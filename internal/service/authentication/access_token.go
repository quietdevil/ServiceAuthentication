package authentication

import (
	"context"
	"github.com/quietdevil/ServiceAuthentication/internal/model"
	"github.com/quietdevil/ServiceAuthentication/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *AuthenticationService) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := utils.VerifyToken(refreshToken, []byte(a.configAuth.RefreshSecretKey()))

	if err != nil {
		return "", err
	}

	accessToken, err := utils.GenerateToken(model.UserInfo{
		Role: claims.Role,
		Name: claims.Username,
	}, []byte(a.configAuth.AccessSecretKey()), a.configAuth.AccessTime())
	if err != nil {
		return "", status.Error(codes.Unauthenticated, "failed to generate token")
	}

	return accessToken, nil
}
