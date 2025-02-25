package authentication

import (
	"context"
	"github.com/quietdevil/ServiceAuthentication/internal/model"
	"github.com/quietdevil/ServiceAuthentication/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *AuthenticationService) GetRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := utils.VerifyToken(refreshToken, []byte(a.configAuth.RefreshSecretKey()))

	if err != nil {
		return "", status.Error(codes.Unauthenticated, "invalid refresh token")
	}

	newRefreshToken, err := utils.GenerateToken(model.UserInfo{
		Role: claims.Role,
		Name: claims.Username,
	}, []byte(a.configAuth.RefreshSecretKey()), a.configAuth.RefreshTime())
	if err != nil {
		return "", status.Error(codes.Unauthenticated, "failed to generate token")
	}
	return newRefreshToken, nil
}
