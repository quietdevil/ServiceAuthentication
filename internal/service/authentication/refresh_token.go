package authentication

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"serviceauth/internal/model"
	"serviceauth/internal/utils"
)

func (a *AuthenticationService) GetRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := utils.VerifyToken(refreshToken, []byte(RefreshSecretKey))

	if err != nil {
		return "", status.Error(codes.Unauthenticated, "invalid refresh token")
	}

	newRefreshToken, err := utils.GenerateToken(model.UserInfo{
		Role: claims.Role,
		Name: claims.Username,
	}, []byte(RefreshSecretKey), refreshTokenExpiration)
	if err != nil {
		return "", status.Error(codes.Unauthenticated, "failed to generate token")
	}
	return newRefreshToken, nil
}
