package authentication

import (
	"context"
	auth "github.com/quietdevil/ServiceAuthentication/pkg/auth_v1"
)

func (i *ImplementationAuthentication) GetRefreshToken(ctx context.Context, req *auth.GetRefreshTokenRequest) (*auth.GetRefreshTokenResponse, error) {
	newRefreshToken, err := i.service.GetRefreshToken(ctx, req.GetOldRefreshToken())
	if err != nil {
		return nil, err
	}
	return &auth.GetRefreshTokenResponse{NewRefreshToken: newRefreshToken}, nil
}
